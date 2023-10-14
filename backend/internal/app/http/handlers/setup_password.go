package handlers

import (
	"github.com/v1shn3vsk7/is-lab/internal/app/http/handlers/adapters"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/v1shn3vsk7/is-lab/internal/models"
)

func (h *Handlers) SetupPassword(c echo.Context) error {
	req := &models.SetupUserPasswordRequest{}
	if err := c.Bind(&req); err != nil {
		return errorWithLog(c, err, "error parse setup user password request", "")
	}

	updateReq, err := adapters.SetupUserPasswordRequestFromAPI(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	if err = h.repo.SetupUserPassword(c.Request().Context(), updateReq); err != nil {
		return errorWithLog(c, err, "setup user password", updateReq)
	}

	return c.JSON(http.StatusOK, &models.EmptyResponse{})
}
