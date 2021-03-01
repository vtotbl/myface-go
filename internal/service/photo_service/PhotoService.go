package photo_service

import (
	"encoding/base64"
	"fmt"
	"image"
	"math/rand"
	"strconv"
	"strings"
	"time"

	_ "image/jpeg"

	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db/photo_repository"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/os/photo_os_repository"
)

type PhotoService struct {
	Repository   *photo_repository.PhotoRepository
	OsRepository *photo_os_repository.PhotoOsRepository
}

func NewPhotoService() (*PhotoService, error) {
	repo, err := photo_repository.NewPhotoRepository()
	if nil != err {
		return nil, err
	}
	service := PhotoService{
		Repository: repo,
	}

	return &service, nil
}

func (service *PhotoService) Upload(userId int, data string) error {
	file, err := service.parsePhoto(data)
	if err != nil {
		return err
	}

	path, err := service.generatePath()
	if err != nil {
		return err
	}

	name, err := service.generateName()
	if err != nil {
		return err
	}

	err = service.OsRepository.UploadJPG(file, path, name)
	if err != nil {
		return err
	}

	photo := domain.Photo{
		Path:   path + name,
		UserId: userId,
	}

	service.Repository.CreatePhoto(photo)

	return nil
}

func (service *PhotoService) CheckCorrectData(data string) error {
	_, err := service.parsePhoto(data)
	if nil != err {
		return err
	}
	return nil
}

func (service *PhotoService) GetById(id int) (string, error) {
	photo, err := service.Repository.GetById(id)
	if nil != err {
		return "", err
	}
	base64, err := service.OsRepository.GetImageBase64(photo.Path)
	if nil != err {
		return "", nil
	}

	return base64, nil
}

func (service *PhotoService) CanGet(userId, photoId int) (bool, error) {
	photo, err := service.Repository.GetById(photoId)
	if nil != err {
		return false, err
	}
	if photo.UserId == userId {
		return true, nil
	}

	return false, nil
}

func (service *PhotoService) GetByUserId(userId int) ([]*PhotoBase64, error) {
	photos, err := service.Repository.GetByUserId(userId)
	if nil != err {
		return nil, err
	}
	if nil == photos {
		return nil, nil
	}

	var base64Photos []*PhotoBase64
	for _, photo := range photos {
		base64, err := service.OsRepository.GetImageBase64(photo.Path)
		if nil != err {
			return nil, err
		}
		base64Photos = append(base64Photos, &PhotoBase64{
			Id:     photo.Id,
			Base64: base64,
		})
	}

	return base64Photos, nil
}

func (service *PhotoService) GetRandom(userId int) (*PhotoBase64, error) {
	photo, err := service.Repository.GetRandom(userId)
	if err != nil {
		return nil, err
	}
	if 0 == photo.Id {
		return nil, nil
	}

	base64, err := service.OsRepository.GetImageBase64(photo.Path)
	if nil != err {
		return nil, err
	}

	return &PhotoBase64{
		Id:     photo.Id,
		Base64: base64,
	}, nil
}

func (service *PhotoService) parsePhoto(data string) (image.Image, error) {
	// если вначале строки есть еще информация, тогда раскоментировать
	// i := strings.Index(data, ",")
	// if i < 0 {
	// 	return nil, errors.New("Invalid format photo")
	// }
	//data = data[i+1:]
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))

	file, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (service *PhotoService) generatePath() (string, error) {
	var pathArr []string
	b := make([]byte, 1)
	var s rand.Source
	var r *rand.Rand
	for i := 1; i < 3; i++ {
		s = rand.NewSource(time.Now().Unix() + int64(i)*time.Hour.Milliseconds())
		r = rand.New(s)

		_, err := r.Read(b)
		if nil != err {
			return "", err
		}

		pathArr = append(pathArr, fmt.Sprintf("%x", b))
	}

	return "docs/images/" + strings.Join(pathArr, "/") + "/", nil
}

func (service *PhotoService) generateName() (string, error) {
	b := make([]byte, 1)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(b)
	if nil != err {
		return "", err
	}

	return fmt.Sprintf("%x", b) + strconv.Itoa(time.Now().Nanosecond()) + ".jpg", nil
}
