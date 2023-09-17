package models

type CreateUserRequest struct {
	Username string `json:"username"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type GetUserRequest struct {
	ID       string `json:"user_id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
