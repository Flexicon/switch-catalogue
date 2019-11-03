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

	games, count, err := h.listingGameService.Feed(page, limit)
	// TODO: investigate some sort of generic error handler
	if err != nil {
		code := http.StatusInternalServerError
		return c.JSON(code, &errorResponse{Message: "Something went wrong", Code: code})
	}

	return c.JSON(http.StatusOK, newGameFeedResponse(games, page, limit, count))
}
