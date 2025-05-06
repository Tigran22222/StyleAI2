package services



import (
	"styleai/internal/models"
"styleai/internal/repositories"
)

type ClothingService struct { clothingRepo *repositories.ClothingRepository outfitRepo *repositories.OutfitRepository }

func NewClothingService(clothingRepo *repositories.ClothingRepository, outfitRepo *repositories.OutfitRepository) *ClothingService { return &ClothingService{ clothingRepo: clothingRepo, outfitRepo: outfitRepo, } }

func (s *ClothingService) GetAll() ([]models.Clothing, error) { return s.clothingRepo.GetAll() }

func (s *ClothingService) CreateOutfit(userID, topID, bottomID, accessoryID int) (models.Outfit, error) { outfit := models.Outfit{ UserID: userID, TopID: topID, BottomID: bottomID, AccessoryID: accessoryID, IsFromWardrobe: false, } return s.outfitRepo.Create(outfit) }
