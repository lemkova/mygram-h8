package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lemkova/mygram-h8/helpers"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
