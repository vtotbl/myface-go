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

	"github.com/Valeriy-Totubalin/myface-go/internal/app/interfaces"
	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
)

type PhotoService struct {
	Repository   interfaces.PhotoDataRepository
	OsRepository interfaces.PhotoFileRepository
}

func NewPhotoService(
	repo interfaces.PhotoDataRepository,
	repoFile interfaces.PhotoFileRepository,
) interfaces.PhotoService {

	return &PhotoService{
		Repository:   repo,
		OsRepository: repoFile,
	}
}

func (service *PhotoService) Upload(userId int, data string) (*domain.Photo, error) {
	file, err := service.parsePhoto(data)
	if err != nil {
		return nil, err
	}

	path, err := service.generatePath()
	if err != nil {
		return nil, err
	}

	name, err := service.generateName()
	if err != nil {
		return nil, err
	}

	err = service.OsRepository.UploadJPG(file, path, name)
	if err != nil {
		return nil, err
	}

	photo := domain.Photo{
		Path:   path + name,
		UserId: userId,
	}

	createdPhoto, err := service.Repository.CreatePhoto(photo)
	if nil != err {
		return nil, err
	}

	return createdPhoto, nil
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

func (service *PhotoService) IsOwner(userId, photoId int) (bool, error) {
	photo, err := service.Repository.GetById(photoId)
	if nil != err {
		return false, err
	}
	if photo.UserId == userId {
		return true, nil
	}

	return false, nil
}

func (service *PhotoService) GetByUserId(userId int) ([]*domain.PhotoBase64, error) {
	photos, err := service.Repository.GetByUserId(userId)
	if nil != err {
		return nil, err
	}
	if nil == photos {
		return nil, nil
	}

	var base64Photos []*domain.PhotoBase64
	for _, photo := range photos {
		base64, err := service.OsRepository.GetImageBase64(photo.Path)
		if nil != err {
			return nil, err
		}
		base64Photos = append(base64Photos, &domain.PhotoBase64{
			Id:     photo.Id,
			Base64: base64,
		})
	}

	return base64Photos, nil
}

func (service *PhotoService) GetRandom(userId int) (*domain.PhotoBase64, error) {
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

	return &domain.PhotoBase64{
		Id:     photo.Id,
		Base64: base64,
	}, nil
}

func (service *PhotoService) Delete(photoId int) error {
	err := service.Repository.Delete(photoId)
	if nil != err {
		return err
	}

	return nil
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
