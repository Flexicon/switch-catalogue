package game

import (
	"github.com/flexicon/switch-catalogue/pkg/store"
)

type Store interface {
	List(offset, limit int) ([]*store.Game, int, error)
}

type ListingService interface {
	Feed(page, limit int) ([]*store.Game, int)
	LastAdded(limit int) ([]*store.Game, int, error)
}

type AddingService interface {
}
