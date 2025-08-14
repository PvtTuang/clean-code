package purchasehistory

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(history *PurchaseHistory) error
	GetByID(id uint64) (*PurchaseHistory, error)
	ListByPlayerID(playerID string) ([]PurchaseHistory, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(history *PurchaseHistory) error {
	return r.db.Create(history).Error
}

func (r *repository) GetByID(id uint64) (*PurchaseHistory, error) {
	var history PurchaseHistory
	if err := r.db.First(&history, id).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func (r *repository) ListByPlayerID(playerID string) ([]PurchaseHistory, error) {
	var histories []PurchaseHistory
	if err := r.db.Where("player_id = ?", playerID).Find(&histories).Error; err != nil {
		return nil, err
	}
	return histories, nil
}
