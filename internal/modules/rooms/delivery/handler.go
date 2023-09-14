package delivery

import (
	authSvc "github.com/doo-dev/pech-pech/internal/modules/auth/usecase"
	"github.com/doo-dev/pech-pech/internal/modules/rooms/presenter"
	"github.com/doo-dev/pech-pech/internal/modules/rooms/usecase"
	"github.com/doo-dev/pech-pech/pkg/httperr"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RoomHandler struct {
	validator presenter.RoomValidator
	roomSvc   usecase.RoomsSvc
}

func NewRoomHandler(roomSvc usecase.RoomsSvc, validator presenter.RoomValidator) RoomHandler {
	return RoomHandler{validator: validator, roomSvc: roomSvc}
}

func (h RoomHandler) CreateRoom(c echo.Context) error {
	var req *presenter.CreateRoomRequest

	fields, err := h.validator.ValidateCreateRoomRequest(req)
	if err != nil {
		code, msg := httperr.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fields,
		})
	}
	userInfo := c.Get("user")
	claims, ok := userInfo.(*authSvc.AuthClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "token is invalid")
	}

	if err := h.roomSvc.CreateRoom(c.Request().Context(), req, claims.ID); err != nil {
		code, msg := httperr.Error(err)
		return c.JSON(code, msg)
	}

	return c.JSON(http.StatusOK, nil)
}

func (h RoomHandler) GetUserRooms(c echo.Context) error {
	userInfo := c.Get("user")
	claims, ok := userInfo.(*authSvc.AuthClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "token is invalid")
	}

	rooms, err := h.roomSvc.GetRooms(c.Request().Context(), claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, rooms)
}
