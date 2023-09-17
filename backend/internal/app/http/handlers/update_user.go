package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/v1shn3vsk7/is-lab/internal/app/http/handlers/adapters"
	"github.com/v1shn3vsk7/is-lab/internal/models"
	modelsRepo "github.com/v1shn3vsk7/is-lab/internal/repository/models"
)

func (h *Handlers) UpdateUserPassword(c echo.Context) error {
	req := &models.UpdateUserPasswordRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	updateReq, err := adapters.UpdateUserPasswordRequestFromAPI(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	user, err := h.repo.GetUser(c.Request().Context(), &modelsRepo.GetUserRequest{
		ID: updateReq.ID,
	})
	if err != nil {
		return errorWithLog(c, err, "get user", "")
	}

	if err = validateUpdateUserPasswordRequest(req, user.IsPasswordConstraint); err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	if err = h.repo.UpdateUserPassword(c.Request().Context(), updateReq); err != nil {
		return errorWithLog(c, err, "update user password", updateReq)
	}

	return c.JSON(http.StatusOK, &models.EmptyResponse{})
}
