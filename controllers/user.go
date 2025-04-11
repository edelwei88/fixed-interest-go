package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/edelwei88/fixed-interest-go/initialize"
	"github.com/edelwei88/fixed-interest-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UsersGET(c *gin.Context) {
	var users []models.User
	result := initialize.DB.Preload("Role").Preload("Loans").Find(&users)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"Users": users,
	})
}

func UserGET(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User
	result := initialize.DB.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
		return
	}
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User": user,
	})
}

func UserPOST(c *gin.Context) {
	var body models.User
	err := c.ShouldBind(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result := initialize.DB.Omit("ID").Create(&body)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User": body,
	})
}

func UserPATCH(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var body models.User
	err = c.ShouldBind(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User
	result := initialize.DB.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
		return
	}
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = initialize.DB.Model(&user).Omit("ID").Updates(&body)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User": user,
	})
}

func UserDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result := initialize.DB.Delete(&models.User{}, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
		return
	}
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
