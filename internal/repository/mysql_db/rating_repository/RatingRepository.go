package rating_repository

import (
	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db"
)

type RatingRepository struct {
}

func NewRatingRepository() (*RatingRepository, error) {
	return &RatingRepository{}, nil
}

func (repo *RatingRepository) CreateRating(rating *domain.Rating) error {
	db, err := mysql_db.GetDB()
	if nil != err {
		return err
	}

	oldRating := Rating{}
	db.Where("user_id = ? AND photo_id = ?", rating.UserId, rating.PhotoId).Find(&oldRating)

	if 0 != oldRating.Id {
		db.Delete(&oldRating)
	}

	db.Create(&Rating{
		Score:   rating.Score,
		PhotoId: rating.PhotoId,
		UserId:  rating.UserId,
	})

	return nil
}

// func (repo *RatingRepository) GetByUserIdPhotoId(userId int, photoId int) (*domain.Rating, error) {
// 	db, err := mysql_db.GetDB()
// 	if nil != err {
// 		return nil, err
// 	}

// 	rating := Rating{}
// 	db.Where("user_id = ? AND photo_id = ?", userId, photoId).Find(&rating)
// 	if 0 == rating.Id {
// 		return nil, errors.New("Rating does not exist")
// 	}

// 	return &domain.Rating{
// 		Id:      rating.Id,
// 		Score:   rating.Score,
// 		PhotoId: rating.PhotoId,
// 		UserId:  rating.UserId,
// 	}, nil
// }
