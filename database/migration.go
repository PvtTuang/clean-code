package database

import (
	"clean-code/pkg/inventory"
	"clean-code/pkg/itemshop"
	"clean-code/pkg/player"
	"clean-code/pkg/playercoin"
	"clean-code/pkg/purchasehistory"

	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&player.Player{},
		&inventory.Inventory{},
		&playercoin.PlayerCoin{},
		&itemshop.Item{},
		&purchasehistory.PurchaseHistory{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migration completed successfully")
}
