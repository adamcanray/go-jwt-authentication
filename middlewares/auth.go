package middlewares

import (
	"go-jwt-authentication/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		authorizationHeader := context.GetHeader("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid access token"})
			context.Abort()
			return
		}
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
