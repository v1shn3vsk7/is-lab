package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/v1shn3vsk7/is-lab/internal/repository/models"
)

func (r *RepoImpl) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.UsersCL.InsertOne(ctx, &user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			// return ErrAlreadyExist
		}
		return fmt.Errorf("error insert user to db, err: %v", err)
	}

	return nil
}
