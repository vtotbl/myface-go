package response

type GetPhoto struct {
	PhotoId string `json:"photo_id" binding:"required"`
	Base64  string `json:"base64" binding:"required,base64"`
}

type Photo struct {
	PhotoId string `json:"photo_id" binding:"required"`
}
