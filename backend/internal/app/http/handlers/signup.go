package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/v1shn3vsk7/is-lab/internal/app/http/handlers/adapters"
	"github.com/v1shn3vsk7/is-lab/internal/models"
)

func (h *Handlers) CreateUser(c echo.Context) error {
	req := &models.CreateUserRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	if err = validateCreateUserRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	if err = h.repo.CreateUser(c.Request().Context(), adapters.UserFromAPI(req)); err != nil {
		return errorWithLog(c, err, "create user", c.Request())
	}

	return c.JSON(http.StatusOK, &models.EmptyResponse{})
}
