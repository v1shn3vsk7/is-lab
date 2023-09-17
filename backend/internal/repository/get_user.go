package repository

import (
	"context"
	"errors"
	"fmt"

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

	return &api.UserToAPI{
		ID:                   user.ID.Hex(),
		Username:             user.Username,
		IsBlocked:            user.Preference.IsBlocked,
		IsPasswordConstraint: user.Preference.IsPasswordConstraint,
	}, nil
}
