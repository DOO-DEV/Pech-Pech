package claims

import (
	"fmt"
	"github.com/doo-dev/pech-pech/internal/modules/auth/usecase"
	"github.com/labstack/echo/v4"
)

func GetClaims(c echo.Context) (*usecase.AuthClaims, error) {
	user := c.Get("user")
	claims, ok := user.(*usecase.AuthClaims)

	if !ok {
		return nil, fmt.Errorf("token is not valid")
	}
	return claims, nil
}
