package repositories

import (
	"styleai/internal/models"
	"gorm.io/gorm"
)

type ClothingRepository struct { db *gorm.DB }

func NewClothingRepository(db *gorm.DB) *ClothingRepository { return &ClothingRepository{db: db} }

func (r *ClothingRepository) GetAll() ([]models.Clothing, error) { var clothing []models.Clothing err := r.db.Find(&clothing).Error return clothing, err }
