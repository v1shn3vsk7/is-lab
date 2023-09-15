package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/v1shn3vsk7/is-lab/internal/config"
)

type RepoImpl struct {
	client  *mongo.Client
	UsersCL *mongo.Collection
}

func New(cfg *config.Config, client *mongo.Client) *RepoImpl {
	return &RepoImpl{
		client:  client,
		UsersCL: client.Database(cfg.MngDBName).Collection("users"),
	}
}
