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

// CreateSocialMedia godoc
// @Summary Create a new social media
// @Description Create a new social media
// @Tags SocialMedia
// @Accept  json, multipart/form-data
// @Produce  json
// @Param name body string true "Social Media Name"
// @Param social_media_url body string true "Social Media URL"
// @Success 201 {object} models.SocialMedia
// @Router /socialmedia/create/ [post]
func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&SocialMedia); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := c.ShouldBind(&SocialMedia); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	SocialMedia.UserID = uint(userData["id"].(float64))

	if err := db.Debug().Create(&SocialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when creating social media",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Social media created successfully",
		"id":      SocialMedia.ID,
	})
}

// GetSocialMedias godoc
// @Summary Get all social media
// @Description Get all social media
// @Tags SocialMedia
// @Produce  json
// @Success 200 {object} []models.SocialMedia
// @Router /socialmedia/ [get]
func GetAllSocialMedias(c *gin.Context) {
	db := database.GetDB()
	var socialMedias []models.SocialMedia
	if err := db.Debug().Find(&socialMedias).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when getting social media",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Social media retrieved successfully",
		"socialMedias": socialMedias,
	})
}

// GetSocialMediasByUserID godoc
// @Summary Get all social media by user id
// @Description Get all social media by user id
// @Tags SocialMedia
// @Produce  json
// @Param userId path int true "User ID"
// @Success 200 {object} []models.SocialMedia
// @Router /socialmedia/user/{userId} [get]
func GetAllSocialMediasByUserID(c *gin.Context) {
	db := database.GetDB()
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid parameter",
			"error":   "bad request",
		})
		return
	}
	var socialMedias []models.SocialMedia
	if err := db.Debug().Where("user_id = ?", userID).Find(&socialMedias).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when getting social media",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Social media retrieved successfully",
		"socialMedias": socialMedias,
	})
}

// GetSocialMedia godoc
// @Summary Get a social media by id
// @Description Get a social media by id
// @Tags SocialMedia
// @Produce  json
// @Param id path int true "Social Media ID"
// @Success 200 {object} models.SocialMedia
// @Router /socialmedia/{id} [get]
func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid parameter",
			"error":   "bad request",
		})
		return
	}
	socialMedia := models.SocialMedia{}
	if err := db.Debug().Where("id = ?", id).First(&socialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when getting social media",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Social media retrieved successfully",
		"socialMedia": socialMedia,
	})
}

// UpdateSocialMedia godoc
// @Summary Update a social media by id
// @Description Update a social media by id
// @Tags SocialMedia
// @Accept  json, multipart/form-data
// @Produce  json
// @Param id path int true "Social Media ID"
// @Param name body string true "Social Media Name"
// @Param social_media_url body string true "Social Media URL"
// @Success 200 {object} models.SocialMedia
// @Router /socialmedia/{id} [put]
func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid parameter",
			"error":   "bad request",
		})
		return
	}
	socialMedia := models.SocialMedia{}
	contentType := helpers.GetContentType(c)

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&socialMedia); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := c.ShouldBind(&socialMedia); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if err := db.Debug().Model(&socialMedia).Where("id = ?", id).Updates(
		models.SocialMedia{
			Name:             socialMedia.Name,
			Social_media_url: socialMedia.Social_media_url,
		}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when updating social media",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Social media updated successfully",
		"socialMedia": socialMedia,
	})
}

// DeleteSocialMedia godoc
// @Summary Delete a social media by id
// @Description Delete a social media by id
// @Tags SocialMedia
// @Produce  json
// @Param id path int true "Social Media ID"
// @Success 200 {object} models.SocialMedia
// @Router /socialmedia/{id} [delete]
func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid parameter",
			"error":   "bad request",
		})
		return
	}
	socialMedia := models.SocialMedia{}
	if err := db.Debug().Where("id = ?", id).Delete(&socialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when deleting social media",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Social media deleted successfully",
	})
}
