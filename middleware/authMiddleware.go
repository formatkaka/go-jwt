package middleware

import (
	"net/http"

	"github.com/formatkaka/go-jwt/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("token")
		if clientToken == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "No Authorization header provided",
			})
			ctx.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)

		if err != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})

			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("first_name", claims.First_name)
		ctx.Set("last_name", claims.Last_name)
		ctx.Set("user_type", claims.User_type)
		ctx.Set("uid", claims.Uid)
		ctx.Next()

	}
}
