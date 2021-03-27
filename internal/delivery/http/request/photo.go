package request

type Upload struct {
	Base64 string `json:"base64" binding:"required,base64"`
}

type Photo struct {
	Id string `uri:"id" binding:"required"`
}

type ChangePhoto struct {
	PhotoId string `json:"photo_id" binding:"required"`
	Base64  string `json:"base64" binding:"required,base64"`
}

type DeletePhoto struct {
	PhotoId int `json:"photo_id" binding:"required"`
}
