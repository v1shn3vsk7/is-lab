package models

type CreateUserRequest struct {
	Username string `json:"username"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CreateUserFromAdminRequest struct {
	Login                string `json:"login"`
	IsPasswordConstraint bool   `json:"is_password_constraint"`
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

type SetupUserPasswordRequest struct {
	ID       string `json:"user_id"`
	Password string `json:"password"`
}
