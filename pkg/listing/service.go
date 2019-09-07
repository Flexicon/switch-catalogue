package listing

import (
	"github.com/flexicon/switch-catalogue/pkg/game"
	"github.com/flexicon/switch-catalogue/pkg/store"
)

type GameService struct {
	store game.Store
}

func NewGameService(store game.Store) *GameService {
	return &GameService{store: store}
}

func (gs *GameService) Feed(page, limit int) ([]*store.Game, int) {
	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	games, count, err := gs.store.List(offset, limit)
	if err != nil {
		return nil, 0
	}

	return games, count
}

func (gs *GameService) LastAdded(limit int) (games []*store.Game, count int, err error) {
	games, count, err = gs.store.List(0, limit)
	if err != nil {
		return nil, 0, err
	}

	return
}
