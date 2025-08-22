package player

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	players := app.Group("/players")
	{
		players.Post("/", h.CreatePlayer)
		players.Get("/:id", h.GetPlayerByID)
		players.Put("/:id", h.UpdatePlayer)
		players.Delete("/:id", h.DeletePlayer)
		players.Get("/", h.ListPlayers)
	}
}

func (h *Handler) CreatePlayer(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetPlayerByID(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) UpdatePlayer(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) DeletePlayer(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) ListPlayers(c *fiber.Ctx) error {
	return nil
}
