
package handlers

import (
	"net/http"
	"strconv"
	"styleai/internal/services"
	"github.com/gin-gonic/gin"
)

type WardrobeHandler struct { wardrobeService *services.WardrobeService }

func NewWardrobeHandler(wardrobeService *services.WardrobeService) *WardrobeHandler { return &WardrobeHandler{wardrobeService: wardrobeService} }

func (h *WardrobeHandler) Add(c *gin.Context) { userID, _ := strconv.Atoi(c.PostForm("user_id")) category := c.PostForm("category") color := c.PostForm("color") file, _, err := c.Request.FormFile("image") if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"}) return } defer file.Close()

	fileBytes := make([]byte, 1024*1024) // Ограничение 1MB
	n, err := file.Read(fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
		return
	}

	userClothing, err := h.wardrobeService.Add(userID, category, color, fileBytes[:n])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Clothing added", "clothing_id": userClothing.ID})

}

func (h *WardrobeHandler) GenerateOutfit(c *gin.Context) { userID, _ := strconv.Atoi(c.Query("user_id")) outfit, err := h.wardrobeService.GenerateOutfit(userID) if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) return } c.JSON(http.StatusOK, gin.H{"outfit": outfit}) }
