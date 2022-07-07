package routes

import (
	controllers "github.com/formatkaka/go-jwt/controllers"
	"github.com/formatkaka/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	routes.Use(middleware.Authenticate())

	routes.GET("/users", controllers.GetUsers())
	routes.GET("/users/:user_id", controllers.GetUser())
}
