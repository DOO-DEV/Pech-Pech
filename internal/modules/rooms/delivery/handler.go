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
	var req presenter.CreateRoomRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

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

	if err := h.roomSvc.CreateRoom(c.Request().Context(), &req, claims.UserID); err != nil {
		code, msg := httperr.Error(err)
		return c.JSON(code, msg)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "room successfully created",
	})
}

func (h RoomHandler) GetUserRooms(c echo.Context) error {
	userInfo := c.Get("user")
	claims, ok := userInfo.(*authSvc.AuthClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "token is invalid")
	}

	rooms, err := h.roomSvc.GetRooms(c.Request().Context(), claims.UserID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, rooms)
}

func (h RoomHandler) DeleteRoom(c echo.Context) error {
	var req presenter.DeleteRoomRequest

	userInfo := c.Get("user")
	claims, ok := userInfo.(*authSvc.AuthClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "token is invalid")
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if field, err := h.validator.ValidateDeleteRequest(req); err != nil {
		code, msg := httperr.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  field,
		})
	}

	if err := h.roomSvc.DeleteRoom(c.Request().Context(), &req, claims.UserID); err != nil {
		code, msg := httperr.Error(err)
		return c.JSON(code, msg)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "room deleted successfully",
	})
}

func (h RoomHandler) UpdateRoom(c echo.Context) error {
	var req presenter.UpdateRoomInfoRequest

	userInfo := c.Get("user")
	claims, ok := userInfo.(*authSvc.AuthClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "token is invalid")
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	field, err := h.validator.ValidateUpdateRoomRequest(req)
	if err != nil {
		code, msg := httperr.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  field,
		})
	}

	room, err := h.roomSvc.UpdateRoomInfo(c.Request().Context(), &req, claims.UserID)
	if err != nil {
		code, msg := httperr.Error(err)
		return c.JSON(code, msg)
	}

	return c.JSON(http.StatusOK, room)
}
