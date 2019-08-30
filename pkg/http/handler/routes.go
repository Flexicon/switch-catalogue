package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterBase(base *echo.Group) {
	base.GET("/ping", h.Ping)
}

func (h *Handler) RegisterApi(api *echo.Group) {
}
