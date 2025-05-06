package handlers

import (
	"net/http"
	"styleai/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StylistTipsHandler struct { db *gorm.DB }

func NewStylistTipsHandler(db *gorm.DB) *StylistTipsHandler { return &StylistTipsHandler{db: db} }

func (h *StylistTipsHandler) GetAll(c *gin.Context) { var tips []models.StylistTip if err := h.db.Find(&tips).Error; err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) return } c.JSON(http.StatusOK, tips) }