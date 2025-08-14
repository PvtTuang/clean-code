package player

import (
	"clean-code/pkg/inventory"
	"time"
)

type Player struct {
	ID          string                `gorm:"primaryKey;type:varchar(64);"`
	Email       string                `gorm:"type:varchar(128);unique;not null;"`
	Name        string                `gorm:"type:varchar(128);not null;"`
	Inventories []inventory.Inventory `gorm:"foreignKey:PlayerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time             `gorm:"not null;autoCreateTime;"`
	UpdatedAt   time.Time             `gorm:"not null;autoUpdateTime;"`
}
