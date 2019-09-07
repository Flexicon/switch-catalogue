package writing

import (
	"github.com/flexicon/switch-catalogue/pkg/game"
	"github.com/flexicon/switch-catalogue/pkg/store"
)

type GameService struct {
	store game.Store
}

func NewGameService(store game.Store) *GameService {
	return &GameService{
		store: store,
	}
}

func (s *GameService) Upsert(game *store.Game) error {
	if game.ID != 0 {
		return s.store.Save(game)
	}

	existing, err := s.store.FindByFsId(game.FsId)
	if err != nil {
		return err
	}

	if existing != nil {
		game.ID = existing.ID
	}

	return s.store.Save(game)
}

func (s *GameService) BatchUpsert(games []*store.Game) error {
	for _, g := range games {
		err := s.Upsert(g)
		if err != nil {
			return err
		}
	}

	return nil
}
