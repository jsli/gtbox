package file

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

func Md5Sum(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Open md5 file failed %s : %s", path, err))
	}
	defer f.Close()

	md5h := md5.New()
	_, err = io.Copy(md5h, f)
	if err != nil {
		return "", errors.New(fmt.Sprintf("copy md5 file failed %s : %s", path, err))
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
		return errors.New(fmt.Sprintf("write failed %s : %s", path, err))
	}
	defer fw.Close()

	_, err = fw.WriteString(content)
	if err != nil {
		return errors.New(fmt.Sprintf("write failed %s : %s", path, err))
	}
	return nil
}
