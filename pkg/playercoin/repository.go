package playercoin

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(coin *PlayerCoin) error
	GetByPlayerID(playerID string) (*PlayerCoin, error)
	Update(coin *PlayerCoin) error
	Delete(id uint64) error
	ListAll() ([]PlayerCoin, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(coin *PlayerCoin) error {
	return r.db.Create(coin).Error
}

func (r *repository) GetByPlayerID(playerID string) (*PlayerCoin, error) {
	var coin PlayerCoin
	result := r.db.First(&coin, "player_id = ?", playerID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &coin, result.Error
}

func (r *repository) Update(coin *PlayerCoin) error {
	return r.db.Save(coin).Error
}

func (r *repository) Delete(id uint64) error {
	return r.db.Delete(&PlayerCoin{}, "id = ?", id).Error
}

func (r *repository) ListAll() ([]PlayerCoin, error) {
	var coins []PlayerCoin
	err := r.db.Find(&coins).Error
	return coins, err
}
