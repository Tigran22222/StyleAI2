package handlers

import (
	"net/http"
	"styleai/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm" )

type ChallengesHandler struct { db *gorm.DB }

func NewChallengesHandler(db *gorm.DB) *ChallengesHandler { return &ChallengesHandler{db: db} }

func (h *ChallengesHandler) GetAll(c *gin.Context) { var challenges []models.Challenge if err := h.db.Find(&challenges).Error; err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) return } c.JSON(http.StatusOK, challenges) }