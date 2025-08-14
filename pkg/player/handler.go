package player

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	players := r.Group("/players")
	{
		players.POST("", h.CreatePlayer)
		players.GET("/:id", h.GetPlayerByID)
		players.PUT("/:id", h.UpdatePlayer)
		players.DELETE("/:id", h.DeletePlayer)
		players.GET("", h.ListPlayers)
	}
}

func (h *Handler) CreatePlayer(c *gin.Context) {
	var player Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreatePlayer(&player); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, player)
}

func (h *Handler) GetPlayerByID(c *gin.Context) {
	id := c.Param("id")

	player, err := h.service.GetPlayerByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if player == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "player not found"})
		return
	}
	c.JSON(http.StatusOK, player)
}

func (h *Handler) UpdatePlayer(c *gin.Context) {
	id := c.Param("id")

	var player Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id != player.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID in path and body do not match"})
		return
	}

	if err := h.service.UpdatePlayer(&player); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

func (h *Handler) DeletePlayer(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeletePlayer(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) ListPlayers(c *gin.Context) {
	players, err := h.service.GetAllPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, players)
}
