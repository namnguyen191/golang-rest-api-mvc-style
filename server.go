package main

import (
	"github.com/gin-gonic/gin"
	"github.com/namnguyen191/themuzix-golang-rest-api/config"
	"github.com/namnguyen191/themuzix-golang-rest-api/controller"
	"github.com/namnguyen191/themuzix-golang-rest-api/repository"
	"github.com/namnguyen191/themuzix-golang-rest-api/service"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	userRepo repository.UserRepository
	authCon  controller.AuthController
	authSer  service.AuthService
	jwtSer   service.JWTService
)

func main() {
	config.SetInitialEnv()

	// connect to the database
	db = config.SetupDatabaseConnection()

	// setup repository
	userRepo = repository.NewUserRepository(db)

	// setup services
	jwtSer = service.NewJWTService()
	authSer = service.NewAuthService(userRepo)

	// setup controller
	authCon = controller.NewAuthController(authSer, jwtSer)

	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authCon.Login)
		authRoutes.POST("/register", authCon.Register)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8081")
}
