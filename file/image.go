package file

import (
	"io"
	"os"
)

const (
	IMAGE_SIZE          = 20971520
	IMAGE_FILE_ACCESS   = 0644
	DEFAULT_BUFFER_SIZE = 4096
	DEFAULT_IMAGE_NAME  = "radio.img"
)

type Component struct {
	Path   string
	Offset int64
}

func GenerateImage(comp_list []Component, dest string, image_size int64) error {
	if dest == "" {
		dest = DEFAULT_IMAGE_NAME
	}

	if image_size <= 0 {
		image_size = IMAGE_SIZE
	}

	if isExist, err := IsFileExist(dest); isExist || err != nil {
		//if it's exist or error occurs when stat it, we should delete it first
		DeleteFile(dest)
	}

	image_file, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, IMAGE_FILE_ACCESS)
	if err != nil {
		return err
	}
	defer image_file.Close()

	buffer := make([]byte, DEFAULT_BUFFER_SIZE)

	for _, src := range comp_list {
		f, err := os.OpenFile(src.Path, os.O_RDONLY, 0)
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
