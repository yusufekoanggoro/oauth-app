package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yusufekoanggoro/oauth-app/backend/config"
	"github.com/yusufekoanggoro/oauth-app/backend/delivery/http/handler"
	"github.com/yusufekoanggoro/oauth-app/backend/repository"
	"github.com/yusufekoanggoro/oauth-app/backend/usecase"
)

func main() {
	cfg := config.LoadConfig()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	authRepo := repository.NewAuthRepository()
	authUsecase := usecase.NewAuthUsecase(authRepo, cfg)
	authHandler := handler.NewAuthHandler(authUsecase)

	r.POST("/auth/google", authHandler.GoogleAuth)

	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
