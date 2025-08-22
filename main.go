package main

import (
	"clean-code/config"
	"clean-code/database"
	"clean-code/pkg/inventory"
	"clean-code/pkg/itemshop"
	"clean-code/pkg/playercoin"
	"clean-code/pkg/purchasehistory"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfigs()

	db := database.Connect(cfg)

	database.RunMigrations(db)

	r := gin.Default()

	// playerRepo := player.NewRepository(db)
	// playerService := player.NewService(playerRepo)
	// playerHandler := player.NewHandler(playerService)
	// playerHandler.RegisterRoutes(r)

	invRepo := inventory.NewRepository(db)
	invService := inventory.NewService(invRepo)
	invHandler := inventory.NewHandler(invService)
	invHandler.RegisterRoutes(r)

	itemRepo := itemshop.NewRepository(db)
	itemService := itemshop.NewService(itemRepo)
	itemHandler := itemshop.NewHandler(itemService)
	itemHandler.RegisterRoutes(r)

	pcRepo := playercoin.NewRepository(db)
	pcService := playercoin.NewService(pcRepo)
	pcHandler := playercoin.NewHandler(pcService)
	pcHandler.RegisterRoutes(r)

	phRepo := purchasehistory.NewRepository(db)
	phService := purchasehistory.NewService(phRepo)
	phHandler := purchasehistory.NewHandler(phService)
	phHandler.RegisterRoutes(r)

	addr := cfg.App.Host + ":" + cfg.App.Port
	log.Printf("Server running at %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
