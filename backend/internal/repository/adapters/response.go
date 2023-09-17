package adapters

import (
	"github.com/v1shn3vsk7/is-lab/internal/models"
	modelsRepo "github.com/v1shn3vsk7/is-lab/internal/repository/models"
)

func UsersToAPI(in []*modelsRepo.User) []*models.UserToAPI {
	res := make([]*models.UserToAPI, 0, len(in))
	for _, val := range in {
		res = append(res, UserToAPI(val))
	}

	return res
}

func UserToAPI(in *modelsRepo.User) *models.UserToAPI {
	return &models.UserToAPI{
		ID:                   in.ID.Hex(),
		Username:             in.Username,
		IsBlocked:            in.Preference.IsBlocked,
		IsPasswordConstraint: in.Preference.IsPasswordConstraint,
	}
}
