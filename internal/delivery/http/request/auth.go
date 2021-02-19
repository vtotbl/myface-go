package request

type SignUp struct {
	Login    string `json:"login" binding:"required,max=60"`
	Password string `json:"password" binding:"required,max=255,min=8"`
	Sex      string `json:"sex" binding:"required,oneof=male female"`
}

type SignIn struct {
	Login    string `json:"login" binding:"required,max=60"`
	Password string `json:"password" binding:"required,max=255,min=8"`
}

type Refresh struct {
	Token string `json:"token" binding:"required,max=64"`
}
