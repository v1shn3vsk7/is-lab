package adapters

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/v1shn3vsk7/is-lab/internal/models"
	modelsRepo "github.com/v1shn3vsk7/is-lab/internal/repository/models"
	"github.com/v1shn3vsk7/is-lab/internal/tech/hash"
)

func UserFromAPI(req *models.CreateUserRequest) *modelsRepo.User {
	res := &modelsRepo.User{
		ID:       primitive.NewObjectID(),
		Username: req.Username,
		Login:    req.Login,
		Preference: &modelsRepo.UserPreference{
			IsBlocked:            false,
			IsPasswordConstraint: false,
		},
	}

	res.Password, res.PasswordSalt = hash.EncryptSecret(req.Password, []byte{})

	return res
}

func GetUserRequestFromAPI(req *models.GetUserRequest) (*modelsRepo.GetUserRequest, error) {
	res := &modelsRepo.GetUserRequest{
		Login: req.Login,
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
	res := &modelsRepo.UpdateUserPasswordRequest{}

	pass, salt := hash.EncryptSecret(req.NewPassword, []byte{})
	res.Password, res.PasswordSalt = pass, salt

	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil, fmt.Errorf("error decode id from api to ObjectId, err: %v", err)
	}
	res.ID = id

	return res, nil
}

func UpdateUserRequestFromAPI(req *models.UpdateUserRequest) (*modelsRepo.UpdateUserRequest, error) {
	res := &modelsRepo.UpdateUserRequest{
		IsBlocked:            req.IsBlocked,
		IsPasswordConstraint: req.IsPasswordConstraint,
	}

	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil, err
	}
	res.ID = id

	return res, nil
}

func SetupUserPasswordRequestFromAPI(req *models.SetupUserPasswordRequest) (*modelsRepo.SetupUserPasswordRequest, error) {
	idHex, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil, err
	}

	passwordHash, salt := hash.EncryptSecret(req.Password, nil)
	return &modelsRepo.SetupUserPasswordRequest{
		ID:           idHex,
		Password:     passwordHash,
		PasswordSalt: salt,
	}, nil
}
