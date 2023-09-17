package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	modelsCmn "github.com/v1shn3vsk7/is-lab/internal/models"
	"github.com/v1shn3vsk7/is-lab/internal/repository/models"
)

func (r *RepoImpl) CreateUser(ctx context.Context, user *models.User) (string, error) {
	res, err := r.UsersCL.InsertOne(ctx, &user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", modelsCmn.ErrAlreadyExists
		}
		return "", fmt.Errorf("error insert user to db, err: %v", err)
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("error obtain user id from db")
	}

	return id.Hex(), nil
}
