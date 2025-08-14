package inventory

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(inv *Inventory) error
	GetByID(id uint64) (*Inventory, error)
	Update(inv *Inventory) error
	Delete(id uint64) error
	ListByPlayerID(playerID string) ([]Inventory, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(inv *Inventory) error {
	return r.db.Create(inv).Error
}

func (r *repository) GetByID(id uint64) (*Inventory, error) {
	var inv Inventory
	result := r.db.First(&inv, "id = ? AND is_delete = false", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &inv, result.Error
}

func (r *repository) Update(inv *Inventory) error {
	return r.db.Save(inv).Error
}

func (r *repository) Delete(id uint64) error {
	return r.db.Model(&Inventory{}).Where("id = ?", id).Update("is_delete", true).Error
}

func (r *repository) ListByPlayerID(playerID string) ([]Inventory, error) {
	var inventories []Inventory
	err := r.db.Where("player_id = ? AND is_dalate = false", playerID).Find(&inventories).Error
	return inventories, err
}
