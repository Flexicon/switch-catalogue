package store

import (
	"github.com/flexicon/switch-catalogue/pkg/model"
	"github.com/jinzhu/gorm"
)

type GameStore struct {
	db *gorm.DB
}

func NewGameStore(db *gorm.DB) *GameStore {
	return &GameStore{db: db}
}

func (s *GameStore) List(offset, limit int) ([]*model.Game, int, error) {
	var (
		games []*model.Game
		count int
	)
	s.db.Model(&model.Game{}).Count(&count)
	s.db.Offset(offset).Limit(limit).Order("created_at desc").Find(&games)

	return games, count, nil
}
