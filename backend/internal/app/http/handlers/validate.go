package handlers

import (
	"errors"
	"regexp"

	"github.com/v1shn3vsk7/is-lab/internal/models"
)

var (
	passwordPattern = regexp.MustCompile(``)
)

func validateCreateUserRequest(req *models.CreateUserRequest) error {
	if len(req.Login) == 0 || len(req.Password) == 0 || len(req.Username) == 0 {
		return errors.New("empty request")
	}

	if !passwordPattern.MatchString(req.Password) {
		return errors.New("password must match pattern")
	}

	return nil
}
