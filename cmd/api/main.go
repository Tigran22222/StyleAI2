package main

import (
	"github.com/gin-gonic/gin"
	"styleai/internal/config"
	"styleai/internal/handlers"
	"styleai/internal/models"
	"styleai/internal/repositories"
	"styleai/internal/services"
)

func main() { // Загрузка конфигурации cfg, err := config.LoadConfig() if err != nil { panic("Failed to load config: " + err.Error()) }

	// Автомиграция таблиц
	cfg.DB.AutoMigrate(
		&models.User{},
		&models.Clothing{},
		&models.UserClothing{},
		&models.Outfit{},
		&models.StylistTip{},
		&models.Challenge{},
		&models.Subscription{},
	)

	// Инициализация репозиториев
	clothingRepo := repositories.NewClothingRepository(cfg.DB)
	wardrobeRepo := repositories.NewWardrobeRepository(cfg.DB, cfg.S3Uploader)
	outfitRepo := repositories.NewOutfitRepository(cfg.DB)
	userRepo := repositories.NewUserRepository(cfg.DB)

	// Инициализация сервисов
	authService := services.NewAuthService(userRepo)
	clothingService := services.NewClothingService(clothingRepo, outfitRepo)
	wardrobeService := services.NewWardrobeService(wardrobeRepo, outfitRepo, userRepo)
	aiStylistService := services.NewAIStylistService(wardrobeRepo)

	// Инициализация обработчиков
	authHandler := handlers.NewAuthHandler(authService)
	clothingHandler := handlers.NewClothingHandler(clothingService)
	wardrobeHandler := handlers.NewWardrobeHandler(wardrobeService)
	aiStylistHandler := handlers.NewAIStylistHandler(aiStylistService)
	stylistTipsHandler := handlers.NewStylistTipsHandler(cfg.DB)
	challengesHandler := handlers.NewChallengesHandler(cfg.DB)

	// Настройка маршрутов
	r := gin.Default()
	r.POST("/auth/sign-up", authHandler.SignUp)
	r.POST("/auth/sign-in", authHandler.SignIn)
	r.GET("/clothing", clothingHandler.GetAll)
	r.POST("/outfit/create", clothingHandler.CreateOutfit)
	r.POST("/wardrobe/add", wardrobeHandler.Add)
	r.POST("/wardrobe/outfit/generate", wardrobeHandler.GenerateOutfit)
	r.POST("/ai-stylist/consult", aiStylistHandler.Consult)
	r.GET("/stylist-tips", stylistTipsHandler.GetAll)
	r.GET("/challenges", challengesHandler.GetAll)

	// Запуск сервера
	r.Run(":8080")

}
