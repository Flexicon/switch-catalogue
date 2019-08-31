package game

import "github.com/flexicon/switch-catalogue/pkg/model"

type Store interface {
	List(offset, limit int) ([]*model.Game, int, error)
}

type ListingService interface {
	Feed(page, limit int) ([]*model.Game, int)
	LastAdded(limit int) ([]*model.Game, int, error)
}

type AddingService interface {
}
