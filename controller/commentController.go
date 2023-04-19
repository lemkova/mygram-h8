package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lemkova/mygram-h8/database"
	"github.com/lemkova/mygram-h8/helpers"
	"github.com/lemkova/mygram-h8/models"
)

// CreateComment godoc
// @Summary Create a new comment
// @Description Create a new comment for a photo
// @Tags Comment
// @Accept  json, multipart/form-data
// @Produce  json
// @Param photoId path int true "Photo ID"
// @Param message body string true "Comment message"
// @Success 201 {object} models.Comment
// @Router /comments/create/{photoId} [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid parameter",
			"error":   "bad request",
		})
		return
	}

	Comment := models.Comment{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&Comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := c.ShouldBind(&Comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	Comment.UserID = uint(userData["id"].(float64))
	Comment.PhotoID = uint(photoId)

	if err := db.Debug().Create(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when creating comment",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment created successfully",
		"comment": Comment,
	})
}

// GetComments godoc
// @Summary Get all comments for a photo
// @Description Get all comments for a photo
// @Tags Comment
// @Produce  json
// @Param photoId path int true "Photo ID"
// @Success 200 {object} models.Comment
// @Router /comments/photo/{photoId} [get]
func GetComments(c *gin.Context) {
	db := database.GetDB()
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid parameter",
			"error":   "bad request",
		})
		return
	}

	Comments := []models.Comment{}
	if err := db.Debug().Where("photo_id = ?", photoId).Find(&Comments).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "photo not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Comments retrieved successfully",
		"comments": Comments,
	})
}

// GetComment godoc
// @Summary Get a comment
// @Description Get a comment by commentId
// @Tags Comment
// @Produce  json
// @Param commentId path int true "Comment ID"
// @Success 200 {object} models.Comment
// @Router /comments/{commentId} [get]
func GetComment(c *gin.Context) {
	db := database.GetDB()
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid parameter",
			"error":   "bad request",
		})
		return
	}

	Comment := models.Comment{}
	if err := db.Debug().Where("id = ?", commentId).Find(&Comment).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "comment not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment retrieved successfully",
		"comment": Comment,
	})
}

// UpdateComment godoc
// @Summary Update a comment
// @Description Update a comment by commentId
// @Tags Comment
// @Accept  json, multipart/form-data
// @Produce  json
// @Param commentId path int true "Comment ID"
// @Param message body string true "Comment Message"
// @Success 200 {object} models.Comment
// @Router /comments/{commentId} [put]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid parameter",
			"error":   "bad request",
		})
		return
	}

	Comment := models.Comment{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&Comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := c.ShouldBind(&Comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if err := db.Debug().Model(&Comment).Where("id = ?", commentId).Updates(
		models.Comment{
			Message: Comment.Message,
		},
	).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when updating comment",
			"error":   err.Error(),
		})
		return
	}
}

// DeleteComment godoc
// @Summary Delete a comment
// @Description Delete a comment by commentId
// @Tags Comment
// @Produce  json
// @Param commentId path int true "Comment ID"
// @Success 200 {object} models.Comment
// @Router /comments/{commentId} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid parameter",
			"error":   "bad request",
		})
		return
	}

	if err := db.Debug().Where("id = ?", commentId).Delete(&models.Comment{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "comment not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted successfully",
	})
}
