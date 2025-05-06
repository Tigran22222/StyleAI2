package handlers

import (
	"net/http"
	"strconv"
	"styleai/internal/models"
	"styleai/internal/services"
	"github.com/gin-gonic/gin"
)

type AIStylistHandler struct { aiStylistService *services.AIStylistService }

func NewAIStylistHandler(aiStylistService *services.AIStylistService) *AIStylistHandler { return &AIStylistHandler{aiStylistService: aiStylistService} }

func (h *AIStylistHandler) Consult(c *gin.Context) { var req models.AIStylistRequest if err := c.BindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"}) return }

	userID, _ := strconv.Atoi(c.Query("user_id"))
	response, err := h.aiStylistService.Consult(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}