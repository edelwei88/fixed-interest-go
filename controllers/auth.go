package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/edelwei88/fixed-interest-go/initialize"
	"github.com/edelwei88/fixed-interest-go/lib"
	"github.com/edelwei88/fixed-interest-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	TokenLength = 128
	UserRole    = "User"
)

func LoginPOST(c *gin.Context) {
	var credentials struct {
		Login    string `binding:"required"`
		Password string `binding:"required"`
	}

	err := c.ShouldBind(&credentials)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var users []models.User
	result := initialize.DB.Where(&models.User{Login: credentials.Login, PasswordHash: lib.HashString(credentials.Password)}).Find(&users)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if len(users) == 0 {
		c.Status(http.StatusUnauthorized)
		return
	}

	bearerToken, err := lib.GenerateBearerToken(TokenLength)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	token := models.Token{
		Token:      bearerToken,
		ExpireDate: time.Now().Add(time.Hour * 24 * 7),
		UserID:     users[0].ID,
	}

	result = initialize.DB.Create(&token)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": token,
		"User":  users[0],
	})
}

func RegisterPOST(c *gin.Context) {
	var credentials struct {
		FirstName   string `binding:"required"`
		LastName    string `binding:"required"`
		PhoneNumber string `binding:"required"`
		Login       string `binding:"required"`
		Password    string `binding:"required"`
	}

	err := c.ShouldBind(&credentials)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var userRole models.Role
	result := initialize.DB.Where(&models.Role{Role: UserRole}).First(&userRole)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	user := models.User{
		FirstName:    credentials.FirstName,
		LastName:     credentials.LastName,
		PhoneNumber:  credentials.PhoneNumber,
		Login:        credentials.Login,
		PasswordHash: lib.HashString(credentials.Password),
		Role:         userRole,
	}

	result = initialize.DB.Create(&user)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	bearerToken, err := lib.GenerateBearerToken(TokenLength)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	token := models.Token{
		Token:      bearerToken,
		ExpireDate: time.Now().Add(time.Hour * 24 * 7),
		UserID:     user.ID,
	}

	result = initialize.DB.Create(&token)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User":  user,
		"Token": token,
	})
}

func CheckBearerTokenPOST(c *gin.Context) {
	var token struct {
		Token string `binding:"required"`
	}

	err := c.ShouldBind(&token)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var existingToken models.Token

	result := initialize.DB.Where(&models.Token{Token: token.Token}).First(&existingToken)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusUnauthorized)
		return
	}
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if existingToken.ExpireDate.Before(time.Now()) {
		c.Status(http.StatusUnauthorized)
		return
	}

	var user models.User
	result = initialize.DB.Where(&models.User{ID: existingToken.UserID}).First(&user)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User":  user,
		"Token": existingToken,
	})
}
