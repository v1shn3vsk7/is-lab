package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/v1shn3vsk7/is-lab/internal/models"
)

func (h *Handlers) Liveness(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.EmptyResponse{})
}

func errorWithLog(c echo.Context, err error, method string, request interface{}) error {
	if errors.Is(err, models.ErrNotFound) {
		return c.JSON(http.StatusNotFound, &models.ErrResponse{Err: err.Error()})
	}

	if errors.Is(err, models.ErrAlreadyExists) {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	log.Err(err).Msgf("unexpected error during %s, err: %v, req: %v", method, err, request)
	return c.JSON(http.StatusInternalServerError, &models.ErrResponse{Err: err.Error()})
}
