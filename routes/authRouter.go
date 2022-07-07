package routes

import (
	controller "github.com/formatkaka/go-jwt/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(routes *gin.Engine) {
	routes.POST("users/signup", controller.SignUp())
	routes.POST("users/login", controller.Login())
}
