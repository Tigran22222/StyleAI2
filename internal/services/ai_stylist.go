package services

import (
	"styleai/internal/models"
	"styleai/internal/repositories"
)

package services

import (
	"styleai/internal/models"
	"styleai/internal/repositories"
)

type AIStylistService struct { wardrobeRepo *repositories.WardrobeRepository }

func NewAIStylistService(wardrobeRepo *repositories.WardrobeRepository) *AIStylistService { return &AIStylistService{wardrobeRepo: wardrobeRepo} }

func (s *AIStylistService) Consult(userID int, req models.AIStylistRequest) (models.AIStylistResponse, error) { wardrobe, err := s.wardrobeRepo.GetByUser(userID) if err != nil { return models.AIStylistResponse{}, err }

	// Простой алгоритм: подобрать вещи из гардероба
	var match models.UserClothing
	for _, item := range wardrobe {
		if item.Category == "bottom" && (req.Color == "red" && item.Color == "black") {
			match = item
			break
		}
	}
	if match.ID == 0 {
		return models.AIStylistResponse{Message: "No matching items found"}, nil
	}

	outfit := models.Outfit{
		UserID:      userID,
		TopID:       0,
		BottomID:    match.ID,
		AccessoryID: 0,
	}
	return models.AIStylistResponse{
		Message: "This red dress pairs well with your black " + match.Category + "!",
		Outfit:  outfit,
	}, nil

}
