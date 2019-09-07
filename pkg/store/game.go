package store

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Game struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Title     string `gorm:"not null"`
}

type GameStore struct {
	db *gorm.DB
}

func NewGameStore(db *gorm.DB) *GameStore {
	return &GameStore{db: db}
}

func (s *GameStore) List(offset, limit int) ([]*Game, int, error) {
	var (
		games []*Game
		count int
	)
	s.db.Model(&Game{}).Count(&count)
	s.db.Offset(offset).Limit(limit).Order("created_at desc").Find(&games)

	return games, count, nil
}

func (s *GameStore) Save(g *Game) error {
	return s.db.Save(g).Error
}
