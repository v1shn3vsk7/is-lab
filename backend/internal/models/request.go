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

type UpdateUserPasswordRequest struct {
	ID          string `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UpdateUserRequest struct {
	ID                   string `json:"user_id"`
	IsPasswordConstraint *bool  `json:"is_password_constraint"`
	IsBlocked            *bool  `json:"is_blocked"`
}
