package repository

import (
	"context"

	"github.com/v1shn3vsk7/is-lab/internal/repository/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
}
