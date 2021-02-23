package response

type GetPhoto struct {
	PhotoId string `json:"photo_id" binding:"required"`
	Photo   string `json:"photo" binding:"required"`
}
