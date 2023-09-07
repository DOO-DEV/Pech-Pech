package delivery

import (
	"github.com/doo-dev/pech-pech/internal/modules/auth/presenter"
	"github.com/doo-dev/pech-pech/internal/modules/auth/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	authSvc usecase.AuthService
}

func NewAuthHandler(authSvc usecase.AuthService) AuthHandler {
	return AuthHandler{authSvc: authSvc}
}

func (h AuthHandler) Register(c echo.Context) error {
	var req presenter.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := h.authSvc.Register(c.Request().Context(), &req)
	if err != nil {
		// TODO - must implement an error richer
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h AuthHandler) Login(c echo.Context) error {
	var req presenter.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := h.authSvc.Login(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
