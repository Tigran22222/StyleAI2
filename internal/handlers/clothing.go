
package handlers

import (
	"net/http"
	"strconv"
	"styleai/internal/models"
	"styleai/internal/services"
	"github.com/gin-gonic/gin"
)

type ClothingHandler struct { clothingService *services.ClothingService }

func NewClothingHandler(clothingService *services.ClothingService) *ClothingHandler { return &ClothingHandler{clothingService: clothingService} }

func (h *ClothingHandler) GetAll(c *gin.Context) { clothing, err := h.clothingService.GetAll() if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) return } c.JSON(http.StatusOK, clothing) }

func (h *ClothingHandler) CreateOutfit(c *gin.Context) { var outfit models.Outfit if err := c.BindJSON(&outfit); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"}) return }

	userID, _ := strconv.Atoi(c.Query("user_id"))
	createdOutfit, err := h.clothingService.CreateOutfit(userID, outfit.TopID, outfit.BottomID, outfit.AccessoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Outfit created", "outfit_id": createdOutfit.ID})

}
