package photo_os_repository

import (
	"encoding/base64"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"io/ioutil"
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

func GetImageBase64(path string) (string, error) {
	pngData, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	b64String := base64.StdEncoding.EncodeToString(pngData)

	return b64String, nil
}
