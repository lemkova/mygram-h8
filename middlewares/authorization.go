package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lemkova/mygram-h8/database"
	"github.com/lemkova/mygram-h8/models"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "invalid parameter",
				"error":   "bad request",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		photo := models.Photo{}

		if err := db.Debug().Select("user_id").First(&photo, uint(photoId)).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "photo not found",
				"error":   err.Error(),
			})
			return
		}

		if photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not authorized to access this photo",
				"error":   "unauthorized",
			})
			return
		}
		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "invalid parameter",
				"error":   "bad request",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		comment := models.Comment{}

		if err := db.Debug().Select("user_id").First(&comment, uint(commentId)).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "comment not found",
				"error":   err.Error(),
			})
			return
		}

		if comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not authorized to access this comment",
				"error":   "unauthorized",
			})
			return
		}
		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialMediaId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "invalid parameter",
				"error":   "bad request",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		socialMedia := models.SocialMedia{}

		if err := db.Debug().Select("user_id").First(&socialMedia, uint(socialMediaId)).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "social media not found",
				"error":   err.Error(),
			})
			return
		}

		if socialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not authorized to access this social media",
				"error":   "unauthorized",
			})
			return
		}
		c.Next()
	}
}
