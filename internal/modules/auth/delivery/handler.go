package delivery

import (
	"github.com/doo-dev/pech-pech/internal/modules/auth/presenter"
	"github.com/doo-dev/pech-pech/internal/modules/auth/usecase"
	"github.com/doo-dev/pech-pech/pkg/httperr"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	authSvc   usecase.AuthService
	validator presenter.Validator
}

func NewAuthHandler(authSvc usecase.AuthService, validator presenter.Validator) AuthHandler {
	return AuthHandler{authSvc: authSvc, validator: validator}
}

func (h AuthHandler) Register(c echo.Context) error {
	req := presenter.RegisterRequest{}

	if err := c.Bind(&req); err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	if fields, err := h.validator.ValidateRegister(req); err != nil {
		code, msg := httperr.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fields,
		})
	}

	res, err := h.authSvc.Register(c.Request().Context(), &req)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h AuthHandler) Login(c echo.Context) error {
	var req presenter.LoginRequest

	if err := c.Bind(&req); err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	if fields, err := h.validator.ValidateLoginRequest(req); err != nil {
		code, msg := httperr.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fields,
		})
	}

	res, err := h.authSvc.Login(c.Request().Context(), &req)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}

func (h AuthHandler) ForgetPassword(c echo.Context) error {
	req := presenter.ForgetPasswodRequest{}

	if err := c.Bind(&req); err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	if fields, err := h.validator.ValidateForgetPasswordRequest(req); err != nil {
		code, msg := httperr.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fields,
		})
	}

	if err := h.authSvc.ForgetPassword(c.Request().Context(), req.Email); err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "verification email successfully send",
	})
}

func (h AuthHandler) ResetPassword(c echo.Context) error {
	req := presenter.ResetPasswordRequest{}

	if err := c.Bind(&req); err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	if fields, err := h.validator.ValidateResetPassword(req); err != nil {
		code, msg := httperr.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fields,
		})
	}

	if err := h.authSvc.ResetPassword(c.Request().Context(), &req); err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "password changed successfully",
	})
}

func (h AuthHandler) UpdatePassword(c echo.Context) error {
	req := presenter.UpdatePasswordRequest{}

	if err := c.Bind(&req); err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	userInfo := c.Get("user")
	claims, ok := userInfo.(*usecase.AuthClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "token is invalid")
	}

	if fields, err := h.validator.ValidateUpdatePasswordRequest(req); err != nil {
		code, msg := httperr.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fields,
		})
	}

	if err := h.authSvc.UpdatePassword(c.Request().Context(), &req, claims.Email); err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"message": "password updated successfully",
	})
}
