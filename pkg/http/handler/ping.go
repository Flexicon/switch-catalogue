package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) Ping(c echo.Context) error {
	type response struct {
		Msg string `json:"msg"`
	}

	return c.JSON(http.StatusOK, response{Msg: "Pong!"})
}
