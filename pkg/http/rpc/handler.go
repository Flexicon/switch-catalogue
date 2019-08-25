package rpc

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"msg": "Pong!"})
}
