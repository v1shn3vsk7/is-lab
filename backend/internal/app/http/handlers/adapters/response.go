package adapters

import (
	api "github.com/v1shn3vsk7/is-lab/internal/models"
	modelsRepo "github.com/v1shn3vsk7/is-lab/internal/repository/models"
)

func UserToAPI(user *modelsRepo.User) *api.UserToAPI {
	res := &api.UserToAPI{
		ID:       user.ID.Hex(),
		Username: user.Username,
	}

	if user.Preference != nil {
		res.IsBlocked = user.Preference.IsBlocked
		res.IsPasswordConstraint = user.Preference.IsPasswordConstraint
	}

	if user.Password == "" {
		res.IsEmptyPassword = true
	}

	return res
}
