package pathutil

import (
	"io/ioutil"
	"os"
	"strings"
)

const (
	SLASH              = string(os.PathSeparator)
	DEFAULT_DIR_ACCESS = 0755
)

func SplitPath(path string) []string {
	return strings.Split(path, SLASH)
}

func ParentPath(path string) string {
	list := SplitPath(path)
	var isAbs bool = false
	if strings.HasPrefix(path, SLASH) {
		list = list[1:]
		isAbs = true
	}

	if strings.HasSuffix(path, SLASH) {
		list = list[:len(list)-1]
	}

	if len(list) <= 0 {
		return SLASH
	} else {
		list = list[:len(list)-1]
		if len(list) <= 0 {
			if isAbs {
				return SLASH
			}
			return ""
		} else {
			parent := strings.Join(list, SLASH)
			if isAbs {
				parent = SLASH + parent
			}
			if parent != "" {
				parent += SLASH
			}
			return parent
		}
	}
}

func BaseName(path string) string {
	list := SplitPath(path)
	if strings.HasSuffix(path, SLASH) {
		list = list[:len(list)-1]
	}
	if list != nil && len(list) > 0 {
		return list[len(list)-1]
	}
	return ""
}

func MkDirSpecificMode(path string, mode os.FileMode) error {
	exist, err := IsExist(path)
	if err == nil {
		if !exist {
			return os.MkdirAll(path, mode)
		} else if exist && err == nil {
			return nil
		}
	}
	return err
}

func MkDir(path string) error {
	exist, err := IsExist(path)
	if err == nil {
		if !exist {
			return os.MkdirAll(path, DEFAULT_DIR_ACCESS)
		} else if exist && err == nil {
			return nil
		}
	}
	return err
}

func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func ListFilesRecursive(prefix, path string, b bool) []string {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}

	list := make([]string, 0, 10)
	var dir_name string
	if !b {
		dir_name = ""
	} else {
		dir_name = BaseName(path) + SLASH
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			tmp_list := ListFilesRecursive(prefix+dir_name, path+info.Name()+SLASH, true)
			list = append(list, tmp_list...)
		} else if info.Mode().IsRegular() {
			list = append(list, prefix+dir_name+info.Name())
		}
	}
	return list
}
