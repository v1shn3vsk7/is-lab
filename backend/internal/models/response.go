package models

type EmptyResponse struct{}

type ErrResponse struct {
	Err string `json:"error"`
}

type CreateUserResponse struct {
	UserID string `json:"user_id"`
}

type UserToAPI struct {
	ID                   string `json:"user_id"`
	Username             string `json:"username"`
	IsBlocked            bool   `json:"is_blocked"`
	IsPasswordConstraint bool   `json:"is_password_constraint"`
	IsEmptyPassword      bool   `json:"is_empty_password"`
	Login                string `json:"login"`
}
