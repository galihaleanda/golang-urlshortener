package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-url-shortener/internal/config"
	"go-url-shortener/internal/handler"
	"go-url-shortener/internal/middleware"
	"go-url-shortener/internal/repository"
	"go-url-shortener/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Database Connection Failed")
	}

	// db.AutoMigrate(&model.URL{})

	repo := repository.NewURLRepository(db)
	service := service.NewURLService(repo)
	handler := handler.NewURLHandler(service)

	r := gin.Default()
	r.Use(middleware.RateLimit())

	r.POST("/shorten", handler.Shorten)
	r.GET("/:code", handler.Redirect)
	r.GET("/analytics/:code", handler.Analytics)

	r.Run(":" + cfg.Port)
}
