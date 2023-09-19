package repository

import (
	"context"

	api "github.com/v1shn3vsk7/is-lab/internal/models"
	"github.com/v1shn3vsk7/is-lab/internal/repository/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
	GetUser(ctx context.Context, req *models.GetUserRequest) (*api.UserToAPI, error)
	GetAllUsers(ctx context.Context) ([]*api.UserToAPI, error)
	UpdateUserPassword(ctx context.Context, req *models.UpdateUserPasswordRequest) error
	UpdateUser(ctx context.Context, req *models.UpdateUserRequest) error
}
