package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"topsis/config"
	"topsis/handler"
	"topsis/internal/infrastructure/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("build")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := gorm.Open(mysql.Open(cfg.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to open database:", err)
	}

	userRepo := repository.NewUserRepository(db)
	standardRepo := repository.NewStandardRepository(db)
	scoreRatingRepo := repository.NewScoreRatingRepository(db)

	initRoute()
}

func initRoute() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := r.Group("v1/api")
	{
		// API users
		api.POST("v1/user", handler.CreateUser)

		// API standards
		api.POST("v1/standards", handler.CreateStandard)
		api.GET("v1/standards", handler.GetStandard)
		api.DELETE("v1/standards", handler.DeleteStandard)

		// API score_ratings
		api.POST("v1/score_ratings", handler.CreateScoreRating)
		api.DELETE("v1/score_ratings", handler.DeleteScoreRating)

		// API Consult
		api.POST("v1/consult", handler.Consult)
	}
	r.Run(":3000")
}
