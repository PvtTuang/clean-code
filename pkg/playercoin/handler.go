package playercoin

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
	coins := r.Group("/playercoins")
	{
		coins.POST("", h.CreateCoin)
		coins.GET("/:playerID", h.GetCoinByPlayerID)
		coins.PUT("/:id", h.UpdateCoin)
		coins.DELETE("/:id", h.DeleteCoin)
		coins.GET("", h.ListCoins)
	}
}

func (h *Handler) CreateCoin(c *gin.Context) {
	var coin PlayerCoin
	if err := c.ShouldBindJSON(&coin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateCoin(&coin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, coin)
}

func (h *Handler) GetCoinByPlayerID(c *gin.Context) {
	playerID := c.Param("playerID")
	coin, err := h.service.GetCoinByPlayerID(playerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if coin == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "coin not found"})
		return
	}
	c.JSON(http.StatusOK, coin)
}

func (h *Handler) UpdateCoin(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var coin PlayerCoin
	if err := c.ShouldBindJSON(&coin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	coin.ID = id

	if err := h.service.UpdateCoin(&coin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coin)
}

func (h *Handler) DeleteCoin(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.service.DeleteCoin(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) ListCoins(c *gin.Context) {
	coins, err := h.service.GetAllCoins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coins)
}
