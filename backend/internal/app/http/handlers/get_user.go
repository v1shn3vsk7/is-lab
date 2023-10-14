package handlers

import (
	"context"
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

	if !validateUserPasswordFromAPI(req, user) {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: "invalid login and/or password"})
	}

	return c.JSON(http.StatusOK, adapters.UserToAPI(user))
}

func (h *Handlers) ListUsers(c echo.Context) error {
	users, err := h.repo.GetAllUsers(context.Background())
	if err != nil {
		return errorWithLog(c, err, "get all users", "")
	}

	return c.JSON(http.StatusOK, users)
}
