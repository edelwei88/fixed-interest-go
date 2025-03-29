package controllers

import (
	"net/http"
	"time"

	"github.com/edelwei88/fixed-interest-go/initialize"
	"github.com/edelwei88/fixed-interest-go/lib"
	"github.com/edelwei88/fixed-interest-go/models"
	"github.com/gin-gonic/gin"
)

const TokenLength = 128

func Login(c *gin.Context) {
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
		"token": token,
		"user":  users[0],
	})
}

func Register(c *gin.Context) {
}

// func CheckBearerToken(c *gin.Context) (models.User, error) {
//
// }
