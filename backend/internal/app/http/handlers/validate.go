package handlers

import (
	"errors"
	"strings"

	"github.com/v1shn3vsk7/is-lab/internal/models"
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
