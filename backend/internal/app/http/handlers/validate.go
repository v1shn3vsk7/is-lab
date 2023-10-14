package handlers

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/v1shn3vsk7/is-lab/internal/models"
	modelsRepo "github.com/v1shn3vsk7/is-lab/internal/repository/models"
	"github.com/v1shn3vsk7/is-lab/internal/tech/hash"
)

func validateCreateUserRequest(req *models.CreateUserRequest) error {
	if len(strings.TrimSpace(req.Login)) == 0 ||
		len(strings.TrimSpace(req.Password)) == 0 ||
		len(strings.TrimSpace(req.Username)) == 0 {
		return errors.New("empty request")
	}

	return nil
}

func validateGetUserRequest(req *models.GetUserRequest) error {
	if strings.TrimSpace(req.ID) == "" &&
		strings.TrimSpace(req.Login) == "" &&
		strings.TrimSpace(req.Password) == "" {
		return errors.New("empty request")
	}

	return nil
}

func validateUpdateUserPasswordRequest(req *models.UpdateUserPasswordRequest) error {
	if strings.TrimSpace(req.ID) == "" &&
		strings.TrimSpace(req.NewPassword) == "" &&
		strings.TrimSpace(req.OldPassword) == "" {
		return errors.New("empty request")
	}

	if strings.TrimSpace(req.ID) == "" {
		return errors.New("user id is empty")
	}

	if strings.TrimSpace(req.NewPassword) == "" {
		return errors.New("new password is empty")
	}

	if strings.TrimSpace(req.OldPassword) == "" {
		return errors.New("olp password is empty")
	}

	return nil
}

func validateUpdateUserRequest(req *models.UpdateUserRequest) error {
	if strings.TrimSpace(req.ID) == "" {
		return errors.New("empty user_id")
	}

	return nil
}

func validateUserPasswordFromAPI(req *models.GetUserRequest, user *modelsRepo.User) bool {
	if req.Password == "" && user.Password == "" {
		return true
	}

	salt, err := base64.StdEncoding.DecodeString(user.PasswordSalt)
	if err != nil {
		return false
	}

	reqPassB, _ := hash.EncryptSecret(req.Password, salt)

	return user.Password == reqPassB
}
