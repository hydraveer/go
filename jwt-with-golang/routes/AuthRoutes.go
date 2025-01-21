package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hydraveer/jwt-with-golang/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
}
