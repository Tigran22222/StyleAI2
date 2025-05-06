package services


import (
	"errors"
	"styleai/internal/models"
	"styleai/internal/repositories"
)

type WardrobeService struct { wardrobeRepo *repositories.WardrobeRepository outfitRepo *repositories.OutfitRepository userRepo *repositories.UserRepository }

func NewWardrobeService(wardrobeRepo *repositories.WardrobeRepository, outfitRepo *repositories.OutfitRepository, userRepo *repositories.UserRepository) *WardrobeService { return &WardrobeService{ wardrobeRepo: wardrobeRepo, outfitRepo: outfitRepo, userRepo: userRepo, } }

func (s *WardrobeService) Add(userID int, category, color string, file []byte) (models.UserClothing, error) { return s.wardrobeRepo.Add(userID, category, color, file) }

func (s *WardrobeService) GenerateOutfit(userID int) (models.Outfit, error) { // Проверка подписки и лимита hasSubscription, err := s.userRepo.HasActiveSubscription(userID) if err != nil { return models.Outfit{}, err } outfitCount, err := s.outfitRepo.CountWardrobeOutfits(userID) if err != nil { return models.Outfit{}, err } if outfitCount >= 5 && !hasSubscription { return models.Outfit{}, errors.New("upgrade to premium for more outfits") }

	// Получение гардероба
	wardrobe, err := s.wardrobeRepo.GetByUser(userID)
	if err != nil {
		return models.Outfit{}, err
	}

	// Простой алгоритм: выбрать топ + низ
	var top, bottom models.UserClothing
	for _, item := range wardrobe {
		if item.Category == "top" {
			top = item
			break
		}
	}
	for _, item := range wardrobe {
		if item.Category == "bottom" {
			bottom = item
			break
		}
	}
	if top.ID == 0 || bottom.ID == 0 {
		return models.Outfit{}, errors.New("not enough items in wardrobe")
	}

	outfit := models.Outfit{
		UserID:         userID,
		TopID:          top.ID,
		BottomID:       bottom.ID,
		AccessoryID:    0,
		IsFromWardrobe: true,
	}
	return s.outfitRepo.Create(outfit)

}
