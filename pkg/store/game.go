package store

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Game struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	Title       string `gorm:"not null"`
	ProductCode string
	FsId        string `gorm:"not null"`
	Url         string
	ChangeDate  time.Time `gorm:"not null"`
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
	err := s.db.Offset(offset).Limit(limit).Order("created_at desc").Find(&games).Error
	if err != nil {
		return nil, 0, err
	}

	return games, count, nil
}

func (s *GameStore) Save(g *Game) error {
	return s.db.Save(g).Error
}

func (s *GameStore) FindByFsId(fsId string) (*Game, error) {
	g := &Game{}
	err := s.db.Where("fs_id = ?", fsId).Last(&g).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		return nil, err
	}

	return g, nil
}
