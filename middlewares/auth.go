package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"org.gfoo/api-rest/utils"
)

func Authenticated(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "Not authorized to access this resource"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.Set("userId", userId)
	context.Next()

}
