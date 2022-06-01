package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"topsis/config"
	"topsis/handler"
	"topsis/internal/domain/usecase"
	"topsis/internal/infrastructure/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("build")
	if err != nil {
		log.Fatal("cannot load config:", err)
		return
	}

	db, err := gorm.Open(mysql.Open(cfg.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to open database:", err)
		return
	}

	userRepo := repository.NewUserRepository(db)
	standardRepo := repository.NewStandardRepository(db)
	scoreRatingRepo := repository.NewScoreRatingRepository(db)

	userDomain := usecase.NewUserDomain(userRepo)
	standardDomain := usecase.NewStandardDomain(standardRepo)
	scoreRatingDomain := usecase.NewScoreRatingDomain(scoreRatingRepo)

	initRoutes(userDomain, standardDomain, scoreRatingDomain)
}

func initRoutes(userDomain *usecase.UserDomain, standardDomain *usecase.StandardDomain, scoreRatingDomain *usecase.ScoreRatingDomain) {
	// init handler
	h := handler.NewHandler(userDomain, standardDomain, scoreRatingDomain)

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
		api.POST("/users", h.CreateUser)

		// API standards
		api.POST("/standards", h.CreateStandard)
		api.GET("/standards", h.GetStandard)
		api.DELETE("/standards", h.DeleteStandard)

		// API score_ratings
		api.POST("/score_ratings", h.CreateScoreRating)
		api.DELETE("/score_ratings", h.DeleteScoreRating)

		// API Consult
		api.POST("/consult", h.Consult)
	}
	err := r.Run(":3000")
	if err != nil {
		return
	}
}
