package game

import (
	"github.com/flexicon/switch-catalogue/pkg/store"
)

type Store interface {
	List(offset, limit int) ([]*store.Game, int, error)
	Save(g *store.Game) error
	FindByFsId(fsId string) (*store.Game, error)
}

type ListingService interface {
	Feed(page, limit int) ([]*store.Game, int, error)
	LastAdded(limit int) ([]*store.Game, int, error)
}

type WritingService interface {
	Upsert(game *store.Game) error
	BatchUpsert(games []*store.Game) error
}
