package model

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Address   string `json:"address"`
	Email     string `json:"email"`
}

type UserRepository interface {
	Create(User)
	SignIn(string, string) bool
}
