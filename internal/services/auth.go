package services

import (
	"errors"
	"styleai/internal/models"
	"styleai/internal/repositories"
	"regexp"
)

type AuthService struct { userRepo *repositories.UserRepository }

func NewAuthService(userRepo *repositories.UserRepository) *AuthService { return &AuthService{userRepo: userRepo} }

func (s *AuthService) SignUp(email, password string) (models.User, error) { // Проверка формата email emailRegex := regexp.MustCompile(^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$) if !emailRegex.MatchString(email) { return models.User{}, errors.New("invalid email format") }

	// Проверка длины пароля
	if len(password) < 6 {
		return models.User{}, errors.New("password must be at least 6 characters")
	}

	user := models.User{
		Email:    email,
		Password: password, // В будущем добавить хэширование пароля
	}
	return s.userRepo.Create(user)

}

func (s *AuthService) SignIn(email, password string) (models.User, error) { user, err := s.userRepo.FindByEmail(email) if err != nil { return models.User{}, err } if user.Password != password { // В будущем сравнивать хэши return models.User{}, errors.New("invalid password") } return user, nil }
