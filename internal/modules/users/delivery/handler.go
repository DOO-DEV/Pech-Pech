package delivery

import (
	"github.com/doo-dev/pech-pech/internal/modules/users/usecase"
	"github.com/doo-dev/pech-pech/pkg/abstract"
	"github.com/doo-dev/pech-pech/pkg/httperr"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userSvc usecase.UserSvc
}

func NewUserHandler(userSvc usecase.UserSvc) UserHandler {
	return UserHandler{userSvc: userSvc}
}

func (h UserHandler) Search(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	size, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	name := c.QueryParam("name")
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name is missing in param")
	}
	pq := abstract.NewPagination(size, page)

	users, err := h.userSvc.SearchUser(c.Request().Context(), name, &pq)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, &users)
}
