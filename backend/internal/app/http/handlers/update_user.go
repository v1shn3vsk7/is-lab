package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/v1shn3vsk7/is-lab/internal/app/http/handlers/adapters"
	"github.com/v1shn3vsk7/is-lab/internal/models"
)

func (h *Handlers) UpdateUser(c echo.Context) error {
	req := &models.UpdateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	if err := validateUpdateUserRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	updateReq, err := adapters.UpdateUserRequestFromAPI(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	if err = h.repo.UpdateUser(c.Request().Context(), updateReq); err != nil {
		return errorWithLog(c, err, "update user", updateReq)
	}

	return c.JSON(http.StatusOK, &models.EmptyResponse{})
}

func (h *Handlers) UpdateUserPassword(c echo.Context) error {
	req := &models.UpdateUserPasswordRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	if err := validateUpdateUserPasswordRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	updateReq, err := adapters.UpdateUserPasswordRequestFromAPI(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	if err = h.repo.UpdateUserPassword(c.Request().Context(), updateReq); err != nil {
		return errorWithLog(c, err, "update user password", updateReq)
	}

	return c.JSON(http.StatusOK, &models.EmptyResponse{})
}
