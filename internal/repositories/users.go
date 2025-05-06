package repositories

import ( "styleai/internal/models" "gorm.io/gorm" )

type UserRepository struct { db *gorm.DB }

func NewUserRepository(db *gorm.DB) *UserRepository { return &UserRepository{db: db} }

func (r *UserRepository) Create(user models.User) (models.User, error) { err := r.db.Create(&user).Error return user, err }

func (r *UserRepository) FindByEmail(email string) (models.User, error) { var user models.User err := r.db.Where("email = ?", email).First(&user).Error return user, err }

func (r *UserRepository) HasActiveSubscription(userID int) (bool, error) { var subscription models.Subscription err := r.db.Where("user_id = ? AND active = ?", userID, true).First(&subscription).Error if err != nil { if err == gorm.ErrRecordNotFound { return false, nil } return false, err } return true, nil }
