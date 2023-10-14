package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"

	modelsCmn "github.com/v1shn3vsk7/is-lab/internal/models"
	modelsRepo "github.com/v1shn3vsk7/is-lab/internal/repository/models"
)

func (r *RepoImpl) UpdateUserPassword(ctx context.Context, req *modelsRepo.UpdateUserPasswordRequest) error {
	filter := FilterByID(req.ID)
	update := bson.M{
		"$set": bson.M{
			"password":      req.Password,
			"password_salt": req.PasswordSalt,
		},
	}

	res, err := r.UsersCL.UpdateOne(ctx, filter, update, nil)
	if err != nil {
		return fmt.Errorf("error updating user password, err: %v", err)
	}

	if res.ModifiedCount == 0 {
		return modelsCmn.ErrAlreadyExists
	}

	return nil
}

func (r *RepoImpl) UpdateUser(ctx context.Context, req *modelsRepo.UpdateUserRequest) error {
	filter := FilterByID(req.ID)
	update := UpdateUserFilter(req)

	res, err := r.UsersCL.UpdateOne(ctx, filter, update, nil)
	if err != nil {
		return fmt.Errorf("error updating user, err: %v", err)
	}

	if res.ModifiedCount == 0 {
		return modelsCmn.ErrAlreadyExists
	}

	return nil
}

func (r *RepoImpl) SetupUserPassword(ctx context.Context, req *modelsRepo.SetupUserPasswordRequest) error {
	filter := FilterByID(req.ID)
	update := bson.M{
		"$set": bson.M{
			"password":      req.Password,
			"password_salt": req.PasswordSalt,
		},
	}

	res, err := r.UsersCL.UpdateOne(ctx, filter, update, nil)
	if err != nil {
		return fmt.Errorf("error update user password, err: %v", err)
	}

	if res.MatchedCount == 0 {
		return modelsCmn.ErrNotFound
	}

	if res.ModifiedCount == 0 {
		return modelsCmn.ErrAlreadyExists
	}

	return nil
}
