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

func (r *RepoImpl) GetUser(ctx context.Context, req *models.GetUserRequest) (res *api.UserToAPI, err error) {
	filter := getUserFilter(req)
	if err = r.UsersCL.FindOne(ctx, filter, nil).Decode(&res); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, modelsCmn.ErrNotFound
		}
		return nil, fmt.Errorf("error find user, err: %v", err)
	}

	return
}
