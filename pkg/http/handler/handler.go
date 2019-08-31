package handler

import "github.com/flexicon/switch-catalogue/pkg/listing"

type Handler struct {
	listingGameService *listing.GameService
}

func NewHandler(lgs *listing.GameService) Handler {
	return Handler{
		listingGameService: lgs,
	}
}
