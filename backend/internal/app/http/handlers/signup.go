package handlers

import (
	modelsRepo "github.com/v1shn3vsk7/is-lab/internal/repository/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/v1shn3vsk7/is-lab/internal/models"
)

func (h *Handlers) CreateUserFromAdmin(c echo.Context) error {
	req := &models.CreateUserFromAdminRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	_, err := h.repo.CreateUser(c.Request().Context(), &modelsRepo.User{
		ID:    primitive.NewObjectID(),
		Login: req.Login,
		Preference: &modelsRepo.UserPreference{
			IsBlocked:            false,
			IsPasswordConstraint: req.IsPasswordConstraint,
		},
	})
	if err != nil {
		return errorWithLog(c, err, "create user", c.Request())
	}

	return c.JSON(http.StatusOK, &models.EmptyResponse{})
}
