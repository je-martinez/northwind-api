package usersRoutes

import (
	usersController "northwind-api/controllers"

	"github.com/gin-gonic/gin"
)

//UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", usersController.SignUp())
	incomingRoutes.POST("/users/login", usersController.Login())
}
