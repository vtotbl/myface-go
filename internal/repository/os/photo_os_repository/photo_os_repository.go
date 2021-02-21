package photo_os_repository

import (
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"os"
)

func UploadJPG(file image.Image, path string, name string) error {
	os.MkdirAll(path, os.ModePerm)
	f, err := os.Create(path + name)
	if err != nil {
		return err
	}
	defer f.Close()

	opt := jpeg.Options{
		Quality: 90,
	}
	err = jpeg.Encode(f, file, &opt)
	if err != nil {
		return err
	}

	return nil
}
