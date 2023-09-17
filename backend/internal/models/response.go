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
	Nickname             string `json:"nickname"`
	IsBlocked            bool   `json:"is_blocked"`
	IsPasswordConstraint bool   `json:"is_password_constraint"`
}
