package repositories

import ( "styleai/internal/models" "styleai/internal/utils" "gorm.io/gorm" )

type WardrobeRepository struct { db *gorm.DB s3Uploader *utils.S3Uploader }

func NewWardrobeRepository(db *gorm.DB, s3Uploader *utils.S3Uploader) *WardrobeRepository { return &WardrobeRepository{db: db, s3Uploader: s3Uploader} }

func (r *WardrobeRepository) Add(userID int, category, color string, file []byte) (models.UserClothing, error) { // Загрузка файла в S3 fileName := "wardrobe/" + string(userID) + "" + category + "" + color + ".jpg" imageURL, err := r.s3Uploader.Upload(fileName, file) if err != nil { return models.UserClothing{}, err }

	userClothing := models.UserClothing{
		UserID:   userID,
		Category: category,
		Color:    color,
		ImageURL: imageURL,
	}
	err = r.db.Create(&userClothing).Error
	return userClothing, err

}

func (r *WardrobeRepository) GetByUser(userID int) ([]models.UserClothing, error) { var wardrobe []models.UserClothing err := r.db.Where("user_id = ?", userID).Find(&wardrobe).Error return wardrobe, err }
