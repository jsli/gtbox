package pathutil

import (
	"strings"
	"os"
)

const (
	SLASH = string(os.PathSeparator)
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
