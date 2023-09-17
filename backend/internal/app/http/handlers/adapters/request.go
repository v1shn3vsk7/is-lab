package adapters

import (
	"fmt"
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

func GetUserRequestFromAPI(req *models.GetUserRequest) (*modelsRepo.GetUserRequest, error) {
	res := &modelsRepo.GetUserRequest{
		Login:    req.Login,
		Password: req.Password,
	}

	if req.ID == "" {
		return res, nil
	}

	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil, err
	}
	res.ID = id

	return res, nil
}

func UpdateUserPasswordRequestFromAPI(req *models.UpdateUserPasswordRequest) (*modelsRepo.UpdateUserPasswordRequest, error) {
	result := &modelsRepo.UpdateUserPasswordRequest{
		Password: req.NewPassword,
	}

	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil, fmt.Errorf("error decode id from api to ObjectId, err: %v", err)
	}
	result.ID = id

	return result, nil
}
