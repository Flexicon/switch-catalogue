package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) GameFeed(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}

	games, count := h.listingGameService.Feed(page, limit)

	return c.JSON(http.StatusOK, newGameFeedResponse(games, page, limit, count))
}
