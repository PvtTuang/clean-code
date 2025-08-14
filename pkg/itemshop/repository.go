package itemshop

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(item *Item) error
	GetByID(id uint64) (*Item, error)
	Update(item *Item) error
	Delete(id uint64) error
	ListAll() ([]Item, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(item *Item) error {
	return r.db.Create(item).Error
}

func (r *repository) GetByID(id uint64) (*Item, error) {
	var item Item
	result := r.db.First(&item, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &item, result.Error
}

func (r *repository) Update(item *Item) error {
	return r.db.Save(item).Error
}

func (r *repository) Delete(id uint64) error {
	return r.db.Delete(&Item{}, "id = ?", id).Error
}

func (r *repository) ListAll() ([]Item, error) {
	var items []Item
	err := r.db.Find(&items).Error
	return items, err
}
