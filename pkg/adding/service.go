package adding

import (
	"github.com/flexicon/switch-catalogue/pkg/game"
)

type GameService struct {
	store game.Store
}

func NewGameService(store game.Store) *GameService {
	return &GameService{
		store: store,
	}
}
