package repository

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	modelsRepo "github.com/v1shn3vsk7/is-lab/internal/repository/models"
)

func (r *RepoImpl) Startup(ctx context.Context) {
	admin := &modelsRepo.User{
		ID:       primitive.NewObjectID(),
		Username: "admin",
		Login:    "admin",
	}

	if _, err := r.UsersCL.InsertOne(ctx, &admin); err != nil {
		if !mongo.IsDuplicateKeyError(err) {
			log.Fatal().Msgf("error creating admin, err: %v", err)
		}
	}
}
