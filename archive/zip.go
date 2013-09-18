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

	for _, zf := range rc.File {
		if strings.HasSuffix(zf.Name, pathutil.SLASH) {
			os.MkdirAll(dest+zf.Name, zf.Mode())
			continue
		}

		reader, err := zf.Open()
		if err != nil {
			return err
		}

		//There is a BUG!!!S
		//can't save ModeSymlink to it !!!!!!
		fw, err := os.OpenFile(dest+pathutil.SLASH+zf.Name, os.O_CREATE|os.O_WRONLY, zf.Mode())
		if err != nil {
			return err
		}

		if _, err := io.Copy(fw, reader); err != nil {
			return err
		}
		fw.Close()
		reader.Close()
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
		f, err := w.Create(str)
		if err != nil {
			return err
		}
		reader, err := os.OpenFile(src+str, os.O_RDONLY, 0644)
		if err != nil {
			return err
		}

		if _, err := io.Copy(f, reader); err != nil {
			return err
		}
		reader.Close()
	}
	return nil
}
