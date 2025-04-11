package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/edelwei88/fixed-interest-go/initialize"
	"github.com/edelwei88/fixed-interest-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddLoanPOST(c *gin.Context) {
	tokenStr, exists := c.Get("bearerToken")
	if !exists {
		c.Status(http.StatusUnauthorized)
		return
	}

	var existingToken models.Token

	result := initialize.DB.Where(&models.Token{Token: tokenStr.(string)}).First(&existingToken)
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

	var loan models.Loan
	loan.UserID = user.ID
	loan.Time = time.Now()
	err := c.ShouldBind(&loan)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result = initialize.DB.Create(&loan)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User": user,
	})
}
