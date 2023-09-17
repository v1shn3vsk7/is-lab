package repository

import (
	"context"
	"fmt"

	modelsCmn "github.com/v1shn3vsk7/is-lab/internal/models"
	modelsRepo "github.com/v1shn3vsk7/is-lab/internal/repository/models"
)

func (r *RepoImpl) UpdateUserPassword(ctx context.Context, req *modelsRepo.UpdateUserPasswordRequest) error {
	filter := FilterByID(req.ID)
	update := UpdateField("password", req.Password)

	res, err := r.UsersCL.UpdateOne(ctx, filter, update, nil)
	if err != nil {
		return fmt.Errorf("error updating user password, err: %v", err)
	}

	if res.ModifiedCount == 0 {
		return modelsCmn.ErrAlreadyExists
	}

	return nil
}
