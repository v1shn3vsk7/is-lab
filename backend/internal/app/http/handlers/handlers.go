package handlers

import "github.com/v1shn3vsk7/is-lab/internal/repository"

type Handlers struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Handlers {
	return &Handlers{
		repo: repo,
	}
}
