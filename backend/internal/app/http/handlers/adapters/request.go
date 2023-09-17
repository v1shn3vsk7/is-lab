package adapters

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/v1shn3vsk7/is-lab/internal/models"
	modelsRepo "github.com/v1shn3vsk7/is-lab/internal/repository/models"
)

func UserFromAPI(req *models.CreateUserRequest) *modelsRepo.User {
	return &modelsRepo.User{
		ID:       primitive.NewObjectID(),
		Username: req.Username,
		Login:    req.Login,
		Password: req.Password,
		Preference: &modelsRepo.UserPreference{
			IsBlocked:            false,
			IsPasswordConstraint: false,
		},
	}
}

func GetUserRequestFromAPI(req *models.GetUserRequest) (res *modelsRepo.GetUserRequest, err error) {
	res.Login = req.Login
	res.Password = req.Password

	if req.ID == "" {
		return
	}

	res.ID, err = primitive.ObjectIDFromHex(req.ID)
	return
}
