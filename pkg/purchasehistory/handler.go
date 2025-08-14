package purchasehistory

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
	routes := r.Group("/purchase-history")
	{
		routes.POST("", h.CreatePurchase)
		routes.GET("/:id", h.GetPurchaseByID)
		routes.GET("/player/:playerID", h.GetPurchaseByPlayerID)
	}
}

func (h *Handler) CreatePurchase(c *gin.Context) {
	var history PurchaseHistory
	if err := c.ShouldBindJSON(&history); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RecordPurchase(&history); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, history)
}

func (h *Handler) GetPurchaseByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	history, err := h.service.GetHistoryByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, history)
}

func (h *Handler) GetPurchaseByPlayerID(c *gin.Context) {
	playerID := c.Param("playerID")

	histories, err := h.service.GetPlayerHistory(playerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, histories)
}
