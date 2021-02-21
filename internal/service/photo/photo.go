package photo

import (
	"encoding/base64"
	"errors"
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

func Upload(userId int, data string) error {
	file, err := parsePhoto(data)
	if err != nil {
		return err
	}

	path, err := generatePath()
	if err != nil {
		return err
	}

	name, err := generateName()
	if err != nil {
		return err
	}

	err = photo_os_repository.UploadJPG(file, path, name)
	if err != nil {
		return err
	}

	photo := domain.Photo{
		Path:   path + name,
		UserId: userId,
	}

	photo_repository.CreatePhoto(photo)

	return nil
}

func CheckCorrectData(data string) error {
	_, err := parsePhoto(data)
	if nil != err {
		return err
	}
	return nil
}

func parsePhoto(data string) (image.Image, error) {
	i := strings.Index(data, ",")
	if i < 0 {
		return nil, errors.New("Invalid format photo")
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[i+1:]))

	file, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func generatePath() (string, error) {
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

func generateName() (string, error) {
	b := make([]byte, 1)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(b)
	if nil != err {
		return "", err
	}

	return fmt.Sprintf("%x", b) + strconv.Itoa(time.Now().Nanosecond()) + ".jpg", nil
}
