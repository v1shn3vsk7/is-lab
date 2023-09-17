package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	
	"github.com/v1shn3vsk7/is-lab/internal/app/http/handlers/adapters"
	"github.com/v1shn3vsk7/is-lab/internal/models"
)

func (h *Handlers) GetUser(c echo.Context) error {
	req := &models.GetUserRequest{
		ID:       c.QueryParam("user_id"),
		Login:    c.QueryParam("login"),
		Password: c.QueryParam("password"),
	}

	if err := validateGetUserRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	getRequest, err := adapters.GetUserRequestFromAPI(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	user, err := h.repo.GetUser(c.Request().Context(), getRequest)
	if err != nil {
		return errorWithLog(c, err, "get user", getRequest)
	}

	return c.JSON(http.StatusOK, &user)
}
