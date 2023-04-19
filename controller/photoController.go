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

// CreatePhoto godoc
// @Summary Create a new photo
// @Description Create a new photo
// @Tags Photo
// @Accept  json, multipart/form-data
// @Produce  json
// @Param photo_title body string true "Photo Title"
// @Param caption body string false "Photo Caption"
// @Param photo_url body string true "Photo URL"
// @Success 201 {object} models.Photo
// @Router /photos/create [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&Photo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := c.ShouldBind(&Photo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	Photo.UserID = uint(userData["id"].(float64))

	if err := db.Debug().Create(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when creating photo",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Photo created successfully",
		"id":      Photo.ID,
	})
}

// GetPhotos godoc
// @Summary Get all photos
// @Description Get all photos from database
// @Tags Photo
// @Produce  json
// @Success 200 {object} []models.Photo
// @Router /photos [get]
func GetPhotos(c *gin.Context) {
	db := database.GetDB()
	Photos := []models.Photo{}

	if err := db.Debug().Find(&Photos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when getting photos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photos retrieved successfully",
		"data":    Photos,
	})
}

// GetPhotosByUser godoc
// @Summary Get all photos by user
// @Description Get all photos by user from database
// @Tags Photo
// @Produce  json
// @Param userId path int true "User ID"
// @Success 200 {object} []models.Photo
// @Router /photos/user/{userId} [get]
func GetPhotosByUser(c *gin.Context) {
	db := database.GetDB()
	Photos := []models.Photo{}
	id := c.Param("userId")

	if err := db.Debug().Where("user_id = ?", id).Find(&Photos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when getting photos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photos retrieved successfully",
		"data":    Photos,
	})
}

// GetPhoto godoc
// @Summary Get a photo
// @Description Get a photo from database
// @Tags Photo
// @Produce  json
// @Param id path int true "Photo ID"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [get]
func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	Photo := models.Photo{}

	if err := db.Debug().First(&Photo, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when getting photo",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo retrieved successfully",
		"data":    Photo,
	})
}

// UpdatePhoto godoc
// @Summary Update a photo
// @Description Update a photo from corresponding ID
// @Tags Photo
// @Accept  json, multipart/form-data
// @Produce  json
// @Param id path int true "Photo ID"
// @Param photo_title body string false "Photo Title"
// @Param caption body string false "Photo Caption"
// @Param photo_url body string false "Photo URL"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [put]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("id"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&Photo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := c.ShouldBind(&Photo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	if err := db.Debug().Model(&Photo).Where("id = ?", photoId).Updates(
		models.Photo{
			Caption:    Photo.Caption,
			PhotoTitle: Photo.PhotoTitle,
			PhotoURL:   Photo.PhotoURL,
		}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when updating photo",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo updated successfully",
	})
}

// DeletePhoto godoc
// @Summary Delete a photo
// @Description Delete a photo from corresponding ID
// @Tags Photo
// @Produce  json
// @Param id path int true "Photo ID"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [delete]
func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}
	photoID, _ := strconv.Atoi(c.Param("id"))

	Photo.ID = uint(photoID)

	if err := db.Debug().Delete(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when deleting photo",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo deleted successfully",
	})
}
