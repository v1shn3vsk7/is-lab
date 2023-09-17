package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/v1shn3vsk7/is-lab/internal/repository/adapters"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	api "github.com/v1shn3vsk7/is-lab/internal/models"
	modelsCmn "github.com/v1shn3vsk7/is-lab/internal/models"
	"github.com/v1shn3vsk7/is-lab/internal/repository/models"
)

func (r *RepoImpl) GetUser(ctx context.Context, req *models.GetUserRequest) (*api.UserToAPI, error) {
	filter := getUserFilter(req)

	var user *models.User
	if err := r.UsersCL.FindOne(ctx, filter, nil).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, modelsCmn.ErrNotFound
		}
		return nil, fmt.Errorf("error find user, err: %v", err)
	}

	return adapters.UserToAPI(user), nil
}

func (r *RepoImpl) GetAllUsers(ctx context.Context) ([]*api.UserToAPI, error) {
	cursor, err := r.UsersCL.Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, fmt.Errorf("error get all users from db, err: %v", err)
	}
	if !cursor.TryNext(ctx) {
		return nil, modelsCmn.ErrNotFound
	}

	var users []*models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, fmt.Errorf("error decode users from db, err: %v", err)
	}

	return adapters.UsersToAPI(users), nil
}
