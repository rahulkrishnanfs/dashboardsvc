package controller

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
