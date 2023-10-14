package app

import (
	"context"

	"github.com/v1shn3vsk7/is-lab/internal/app/http"
	"github.com/v1shn3vsk7/is-lab/internal/bootstrap"
	"github.com/v1shn3vsk7/is-lab/internal/config"
	"github.com/v1shn3vsk7/is-lab/internal/repository"
	"github.com/v1shn3vsk7/is-lab/internal/tech/hash"
	"github.com/v1shn3vsk7/is-lab/internal/utils"
)

func Run(ctx context.Context, cfg *config.Config) error {
	_, cancel := context.WithCancel(ctx)

	hash.NewSecretKey(cfg.SecretKey)

	mng := bootstrap.New(ctx, cfg)

	repo := repository.New(cfg, mng)
	repo.CreateIndexes(ctx)
	repo.Startup(ctx)

	httpServer := http.New(cfg, repo)
	httpServer.Start()

	utils.GracefulShutdown(cancel)

	return nil
}
