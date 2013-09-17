package file

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/jsli/gtbox/pathutil"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func Md5Sum(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("Open md5 file failed %s : %s", path, err)
	}
	defer f.Close()

	md5h := md5.New()
	_, err = io.Copy(md5h, f)
	if err != nil {
		return "", fmt.Errorf("copy md5 file failed %s : %s", path, err)
	}

	md5_str := hex.EncodeToString(md5h.Sum(nil))
	return md5_str, nil
}

func IsFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func DeleteFile(path string) error {
	return do_delete(path)
}

func DeleteDir(path string) error {
	return do_delete(path)
}

func do_delete(path string) error {
	return os.RemoveAll(path)
}

func WriteString2File(content, path string) error {
	fw, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("write failed %s : %s", path, err)
	}
	defer fw.Close()

	_, err = fw.WriteString(content)
	if err != nil {
		return fmt.Errorf("write failed %s : %s", path, err)
	}
	return nil
}

func CopyFile(src, dest string) (int64, error) {
	fi, fi_err := os.Stat(src)
	if fi_err != nil {
		return 0, fi_err
	}
	srcFile, err := os.OpenFile(src, os.O_RDONLY, fi.Mode())
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	if strings.HasSuffix(dest, pathutil.SLASH) {
		pathutil.MkDir(dest)
		dest += pathutil.BaseName(src)
	} else {
		DeleteFile(dest)
		parent_path := pathutil.ParentPath(dest)
		pathutil.MkDir(parent_path)
	}

	destFile, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, fi.Mode())
	if err != nil {
		return 0, err
	}
	defer destFile.Close()

	return io.Copy(destFile, srcFile)
}

func CopyDir(src, dest string) error {
	DeleteDir(dest)
	mode, err := GetFileMode(src)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(dest, pathutil.SLASH) {
		dest += pathutil.SLASH
		err := pathutil.MkDirSpecificMode(dest, mode)
		if err != nil {
			return err
		}
	}

	if !strings.HasSuffix(src, pathutil.SLASH) {
		src += pathutil.SLASH
	}

	fileInfos, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			err := CopyDir(src+info.Name()+pathutil.SLASH, dest+info.Name()+pathutil.SLASH)
			if err != nil {
				return err
			}
		} else if info.Mode().IsRegular() {
			_, err := CopyFile(src+info.Name(), dest+info.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetFileMode(path string) (os.FileMode, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return fi.Mode(), nil
}

//func CopyFile(src io.Reader, dest string) error {
//	tag := "CopyFile"
//	content, err := ioutil.ReadAll(src)
//	if err != nil || len(content) == 0 {
//		log.Log(tag, fmt.Sprintf("Failed to read file:\n err = %s", err))
//		return err
//	}
//
//	destFile, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, constant.DEFAULT_FILE_ACCESS)
//	if err != nil {
//		log.Log(tag, fmt.Sprintf("Failed to open file %s :\n err = %s", dest, err))
//		return err
//	}
//	defer destFile.Close()
//
//	_, err = destFile.Write(content)
//	if err != nil {
//		log.Log(tag, fmt.Sprintf("Failed to write file %s :\n err = %s", dest, err))
//		return err
//	}
//	log.Log(tag, fmt.Sprintf("Copy to %s success!", dest))
//	return nil
//}