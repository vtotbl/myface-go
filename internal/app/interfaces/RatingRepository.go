package interfaces

import "github.com/Valeriy-Totubalin/myface-go/internal/domain"

type RatingRepository interface {
	CreateRating(rating *domain.Rating) error
}
