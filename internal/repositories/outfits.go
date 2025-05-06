package repositories



import ( "styleai/internal/models" "gorm.io/gorm" )

type OutfitRepository struct { db *gorm.DB }

func NewOutfitRepository(db *gorm.DB) *OutfitRepository { return &OutfitRepository{db: db} }

func (r *OutfitRepository) Create(outfit models.Outfit) (models.Outfit, error) { err := r.db.Create(&outfit).Error return outfit, err }

func (r *OutfitRepository) CountWardrobeOutfits(userID int) (int64, error) { var count int64 err := r.db.Model(&models.Outfit{}).Where("user_id = ? AND is_from_wardrobe = ?", userID, true).Count(&count).Error return count, err }
