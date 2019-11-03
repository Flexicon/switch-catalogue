package handler

import (
	"github.com/flexicon/switch-catalogue/pkg/store"
	"math"
)

type gameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type gameFeedResponse struct {
	Games      []*gameResponse `json:"games"`
	Page       int             `json:"page"`
	TotalPages int             `json:"total_pages"`
	Limit      int             `json:"limit"`
	Count      int             `json:"count"`
}

type errorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func newGameResponse(game *store.Game) *gameResponse {
	return &gameResponse{
		ID:    game.ID,
		Title: game.Title,
	}
}

func newGameFeedResponse(games []*store.Game, page, limit, count int) *gameFeedResponse {
	r := &gameFeedResponse{
		Games:      make([]*gameResponse, 0),
		Page:       page,
		Limit:      limit,
		Count:      count,
		TotalPages: int(math.Ceil(float64(count) / float64(limit))),
	}

	for _, game := range games {
		r.Games = append(r.Games, newGameResponse(game))
	}

	return r
}
