package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) Liveness(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
