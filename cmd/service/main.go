package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"topsis/config"

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
		api.POST("v1/user", controller.CreateUser)

		// API standards
		api.POST("v1/standards", controller.CountMessageNoRead)
		api.GET("v1/standards")
		api.DELETE("v1/standards")

		// API score_ratings
		api.POST("v1/score_ratings")
		api.DELETE("v1/score_ratings")

		// API Consult
		api.POST("v1/consult")
	}
	r.Run(":3000")
}
