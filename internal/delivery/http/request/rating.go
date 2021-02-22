package request

type SetRatingInput struct {
	Rating  float64 `json:"rating" binding:"required,min=1,max=10"`
	PhotoId int     `json:"photo_id" binding:"required"`
}
