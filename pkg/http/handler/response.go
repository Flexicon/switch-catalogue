package handler

import (
	"github.com/flexicon/switch-catalogue/pkg/model"
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

func newGameResponse(game *model.Game) *gameResponse {
	return &gameResponse{
		ID:    game.ID,
		Title: game.Title,
	}
}

func newGameFeedResponse(games []*model.Game, page, limit, count int) *gameFeedResponse {
	r := &gameFeedResponse{}
	r.Games = make([]*gameResponse, 0)
	r.Page = page
	r.Limit = limit
	r.Count = count
	r.TotalPages = int(math.Ceil(float64(count) / float64(limit)))

	for _, game := range games {
		r.Games = append(r.Games, newGameResponse(game))
	}

	return r
}
