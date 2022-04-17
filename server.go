package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salamidrus/config"
	"github.com/salamidrus/controller"

	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
