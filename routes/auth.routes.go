package usersRoutes

import (
	authController "northwind-api/controllers/auth"

	"github.com/gin-gonic/gin"
)

//UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/auth/signup", authController.SignUp())
	incomingRoutes.POST("/auth/login", authController.Login())
	incomingRoutes.POST("/auth/refresh-token", authController.RefreshToken())
}
