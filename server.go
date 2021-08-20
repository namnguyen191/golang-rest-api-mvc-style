package main

import (
	"github.com/gin-gonic/gin"
	"github.com/namnguyen191/themuzix-golang-rest-api/config"
	"github.com/namnguyen191/themuzix-golang-rest-api/controller"
	"github.com/namnguyen191/themuzix-golang-rest-api/entity"
	"github.com/namnguyen191/themuzix-golang-rest-api/middleware"
	"github.com/namnguyen191/themuzix-golang-rest-api/repository"
	"github.com/namnguyen191/themuzix-golang-rest-api/service"
)

func main() {
	config.SetInitialEnv()

	// Connect to the database
	db := config.SetupDatabaseConnection()
	// Migrate DB model
	err := db.AutoMigrate(&entity.Artist{})
	if err != nil {
		panic("fail to migrate Artist table")
	}
	err = db.AutoMigrate(&entity.Song{})
	if err != nil {
		panic("fail to migrate Song table")
	}
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic("fail to migrate User table")
	}

	r := gin.Default()

	authController := controller.NewAuthController(
		service.NewAuthService(repository.NewUserRepository(db)),
		service.NewJWTService(),
	)

	authRoutes := r.Group("api/auth", middleware.AuthorizeJWT(service.NewJWTService()))
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8081")
}
