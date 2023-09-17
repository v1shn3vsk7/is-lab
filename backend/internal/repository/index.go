package repository

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *RepoImpl) CreateIndexes(ctx context.Context) {
	loginIdx := mongo.IndexModel{
		Keys:    bson.D{{Key: "login", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	if _, err := r.UsersCL.Indexes().CreateOne(ctx, loginIdx, nil); err != nil {
		log.Fatal().Msgf("error create index for usersCL, err: %v", err)
	}
}
