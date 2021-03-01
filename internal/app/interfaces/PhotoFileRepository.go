package interfaces

import "image"

type PhotoFileRepository interface {
	UploadJPG(file image.Image, path string, name string) error
	GetImageBase64(path string) (string, error)
}
