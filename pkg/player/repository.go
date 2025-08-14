package player

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(player *Player) error
	GetByID(id string) (*Player, error)
	Update(player *Player) error
	Delete(id string) error
	ListAll() ([]Player, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(player *Player) error {
	return r.db.Create(player).Error
}

func (r *repository) GetByID(id string) (*Player, error) {
	var player Player
	result := r.db.Preload("Inventories").First(&player, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &player, result.Error
}

func (r *repository) Update(player *Player) error {
	return r.db.Save(player).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&Player{}, "id = ?", id).Error
}

func (r *repository) ListAll() ([]Player, error) {
	var players []Player
	err := r.db.Preload("Inventories").Find(&players).Error
	return players, err
}
