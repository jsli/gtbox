package ota

import (
	"github.com/jsli/gtbox/sys"
	"github.com/jsli/gtbox/file"
	"github.com/jsli/gtbox/pathutil"
	"errors"
	"fmt"
)

var (
	DEBUG bool = true
)

func GenerateOtaPackage(cmd string, params []string) error {
	res, output, err := sys.ExecCmd(cmd, params)
	if DEBUG {
		fmt.Println(output)
	}
	if !res || err != nil {
		return errors.New(fmt.Sprintf("%s failed: %s\n\tdetail message: %s\n", cmd, err, output))
	}
	return nil
}

func RecordMd5(path, txt_path string) error {
	md5_str, err := file.Md5Sum(path)
	if md5_str == "" || err != nil{
		return err
	} else {
		parent := pathutil.ParentPath(path)
		if parent == "" {
			return errors.New(fmt.Sprintf("record md5 failed: %s parent path is empty", path))
		}
		base_name := pathutil.BaseName(path)
		err := file.WriteString2File(fmt.Sprintf("%s %s", md5_str, base_name), txt_path)
		if err != nil {
			return err
		}
	}
	return nil
}