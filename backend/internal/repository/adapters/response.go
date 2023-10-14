package adapters

import (
	"github.com/v1shn3vsk7/is-lab/internal/models"
	modelsRepo "github.com/v1shn3vsk7/is-lab/internal/repository/models"
)

func UsersToAPI(in []*modelsRepo.User) []*models.UserToAPI {
	res := make([]*models.UserToAPI, 0, len(in))
	for _, val := range in {
		if val.Login == "admin" {
			continue
		}
		res = append(res, UserToAPI(val))
	}

	return res
}

func UserToAPI(in *modelsRepo.User) *models.UserToAPI {
	res := &models.UserToAPI{
		ID:       in.ID.Hex(),
		Username: in.Username,
	}

	if in.Preference != nil {
		res.IsBlocked = in.Preference.IsBlocked
		res.IsPasswordConstraint = in.Preference.IsPasswordConstraint
	}

	res.Login = in.Login

	return res
}
