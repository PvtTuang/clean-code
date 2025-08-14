package inventory

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	inv := r.Group("/inventories")
	{
		inv.POST("", h.AddInventory)
		inv.GET("/:id", h.GetInventoryByID)
		inv.PUT("/:id", h.UpdateInventory)
		inv.DELETE("/:id", h.RemoveInventory)
		inv.GET("/player/:playerID", h.GetInventoriesByPlayer)
	}
}

func (h *Handler) AddInventory(c *gin.Context) {
	var inv Inventory
	if err := c.ShouldBindJSON(&inv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AddInventory(&inv); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, inv)
}

func (h *Handler) GetInventoryByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	inv, err := h.service.GetInventoryByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if inv == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "inventory not found"})
		return
	}

	c.JSON(http.StatusOK, inv)
}

func (h *Handler) UpdateInventory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var inv Inventory
	if err := c.ShouldBindJSON(&inv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if inv.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id in path and body do not match"})
		return
	}

	if err := h.service.UpdateInventory(&inv); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, inv)
}

func (h *Handler) RemoveInventory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.RemoveInventory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) GetInventoriesByPlayer(c *gin.Context) {
	playerID := c.Param("playerID")
	inventories, err := h.service.GetInventoriesByPlayer(playerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, inventories)
}
