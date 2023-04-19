package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lemkova/mygram-h8/database"
	"github.com/lemkova/mygram-h8/helpers"
	"github.com/lemkova/mygram-h8/models"
)

var (
	appJSON = "application/json"
)

// UserRegister godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags User
// @Accept  json, multipart/form-data
// @Produce  json
// @Param username body string true "Username"
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Param age body int true "Age"
// @Success 201 {object} models.User
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if err := db.Debug().Create(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed when registering user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "User registered successfully",
		"id":       User.ID,
		"username": User.Username,
		"email":    User.Email,
	})
}

// UserLogin godoc
// @Summary Login a user
// @Description Login a user
// @Tags User
// @Accept  json, multipart/form-data
// @Produce  json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} models.User
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := c.ShouldBind(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	password := User.Password

	if err := db.Where("email = ?", User.Email).Take(&User).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed when logging in user",
			"error":   "Email or password is incorrect",
		})
		return
	}

	if !helpers.ComparePass([]byte(User.Password), []byte(password)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed when logging in user",
			"error":   "Email or password is incorrect",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"token":   token,
	})
}
