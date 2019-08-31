package listing

import (
	"github.com/flexicon/switch-catalogue/pkg/model"
	"github.com/flexicon/switch-catalogue/pkg/store"
)

type GameService struct {
	store *store.GameStore
}

func NewGameService(store *store.GameStore) *GameService {
	return &GameService{store: store}
}

func (gs *GameService) Feed(page, limit int) ([]*model.Game, int) {
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
