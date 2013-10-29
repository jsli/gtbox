package archive

import (
	"archive/zip"
	"github.com/jsli/gtbox/pathutil"
	"io"
	"os"
	"strings"
)

func ExtractZipFile(src string, dest string) error {
	if !strings.HasSuffix(dest, pathutil.SLASH) {
		dest += pathutil.SLASH
	}
	rc, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer rc.Close()

	for _, zf := range rc.Reader.File {
		name := zf.Name
		mode := zf.Mode()
		if mode.IsDir() {
			os.MkdirAll(dest+name, zf.Mode())
		} else {
			parent_path := pathutil.ParentPath(dest + name)
			os.MkdirAll(parent_path, pathutil.DEFAULT_DIR_ACCESS)
			if err := unpackZippedFile(dest, name, zf); err != nil {
				return err
			}
			os.Chmod(dest+name, mode)
		}
	}

	return nil
}

func unpackZippedFile(path string, name string, zf *zip.File) error {
	writer, err := os.Create(path + name)
	if err != nil {
		return err
	}
	defer writer.Close()

	reader, err := zf.Open()
	if err != nil {
		return err
	}
	defer reader.Close()

	if _, err := io.Copy(writer, reader); err != nil {
		return err
	}

	return nil
}

func ArchiveZipFile(src, dest string) error {
	list := pathutil.ListFilesRecursive("", src, false)
	zw, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer zw.Close()

	w := zip.NewWriter(zw)
	if w == nil {
		return err
	}
	defer w.Close()

	for _, str := range list {
		if err := writeFile2Zip(w, src, str); err != nil {
			return err
		}
	}
	return nil
}

func writeFile2Zip(zw *zip.Writer, path string, name string) error {
	file, err := os.Open(path + name)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = name

	writer, err := zw.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}
