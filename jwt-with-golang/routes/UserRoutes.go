package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hydraveer/jwt-with-golang/controllers"
	middlewares "github.com/hydraveer/jwt-with-golang/middlewares"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.GET("/users",controllers.GetUsers())
	incomingRoutes.GET("/user/:user_id",controllers.GetUser())
}