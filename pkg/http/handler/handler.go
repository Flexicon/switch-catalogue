package handler

import (
	"github.com/flexicon/switch-catalogue/pkg/game"
)

type Handler struct {
	listingGameService game.ListingService
}

func NewHandler(lgs game.ListingService) Handler {
	return Handler{
		listingGameService: lgs,
	}
}
