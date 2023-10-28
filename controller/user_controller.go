package controller

import (
	"dashboardsvc/auth"
	"dashboardsvc/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserController struct {
	repository model.UserRepository
}

func NewUserController(repository model.UserRepository) *UserController {
	return &UserController{
		repository: repository,
	}
}

func (u UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user model.User
	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	u.repository.Create(user)
}

func (u UserController) SignIn(w http.ResponseWriter, r *http.Request) {
	var userLogin UserLogin
	json.NewDecoder(r.Body).Decode(&userLogin)
	fmt.Println(userLogin.Username, userLogin.Password, "body")
	check := u.repository.SignIn(userLogin.Username, userLogin.Password)
	if check {
		fmt.Println("Login succeded")
		token, _ := auth.GenerateToken(userLogin.Username)
		refreshToken, _ := auth.GenerateRefreshToken(userLogin.Username)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusCreated)
		j, err := json.Marshal(UserToken{AccessToken: token, RefreshToken: refreshToken})
		if err != nil {
			fmt.Println("unexpected error")
			return
		}
		w.Write(j)
	}
}
