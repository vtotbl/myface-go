package request

type Upload struct {
	Photo string `json:"photo" binding:"required,base64"`
}

type Photo struct {
	Id string `uri:"id" binding:"required"`
}
