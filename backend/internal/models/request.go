package models

type CreateUserRequest struct {
	Username string `json:"username"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
