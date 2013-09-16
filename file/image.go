package file

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	IMAGE_SIZE          = 20971520
	ROOT_DIR_KEY        = "root_dir"
	IMAGE_SIZE_KEY      = "image_size"
	DEST_PATH_KEY       = "dest_path"
	IMAGE_FILE_ACCESS   = 0644
	DEFAULT_BUFFER_SIZE = 4096
)

type Component struct {
	Name   string
	Offset int64
}

func GenerateImage(comp_list []Component, args map[string]interface{}) error {
	var image_size int64 = IMAGE_SIZE
	var root_dir string
	var dest_path string

	//we should check it's type???It must be a map here.
	if args == nil {
		return errors.New(fmt.Sprintf("args cannot be nil!\n"))
	}

	if _, ok := args[ROOT_DIR_KEY]; !ok {
		return errors.New(fmt.Sprintf("args miss %s!\n", ROOT_DIR_KEY))
	}
	root_dir = args[ROOT_DIR_KEY].(string)

	if _, ok := args[DEST_PATH_KEY]; !ok {
		return errors.New(fmt.Sprintf("args miss %s!\n", DEST_PATH_KEY))
	}
	dest_path = args[DEST_PATH_KEY].(string)

	if _, ok := args[IMAGE_SIZE_KEY]; ok {
		image_size = args[IMAGE_SIZE_KEY].(int64)
	}
	dest_path = args[DEST_PATH_KEY].(string)

	if isExist, err := IsFileExist(dest_path); isExist || err != nil {
		//if it's exist or error occurs when stat it, we should delete it first
		DeleteFile(dest_path)
	}

	image_file, err := os.OpenFile(dest_path, os.O_WRONLY|os.O_CREATE, IMAGE_FILE_ACCESS)
	if err != nil {
		return err
	}
	defer image_file.Close()

	buffer := make([]byte, DEFAULT_BUFFER_SIZE)

	for _, src := range comp_list {
		f, err := os.OpenFile(root_dir+src.Name, os.O_RDONLY, 0)
		if err != nil {
			return err
		}
		defer f.Close() // dup close opt, in case abnormally return in the for-cycle below

		image_file.Seek(src.Offset, 0)
		for {
			n, err := f.Read(buffer)
			if n > 0 && err == nil {
				n, err = image_file.Write(buffer[:n])
				if err != nil {
					return err
				}
				continue
			} else if err == io.EOF {
				break
			} else if err != nil {
				return err
			}
		}
		f.Close()
	}
	image_file.Truncate(image_size)
	return nil
}
