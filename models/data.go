package models

type LoginRequest struct {
	UserName string `json:"username" form:"username" xml:"username" query:"`
	Password string `json:"password" form:"password" xml:"password" query:"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type X struct {
	Text string `json:text`
}