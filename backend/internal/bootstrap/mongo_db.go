package bootstrap

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/v1shn3vsk7/is-lab/internal/config"
	"github.com/v1shn3vsk7/is-lab/internal/tech/closer"
)

func New(ctx context.Context, cfg *config.Config) *mongo.Client {
	opts := options.Client().ApplyURI(cfg.MngDSN)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal().Msgf("error connect to mongo, err: %v", err)
	}

	go func() {
		t := time.NewTicker(cfg.MngPingInterval)
		for range t.C {
			if err = client.Ping(ctx, nil); err != nil {
				log.Error().Msgf("error mng ping, err: %v", err)
			}
		}
	}()

	closer.Add(func() error {
		return client.Disconnect(context.TODO())
	})

	return client
}
