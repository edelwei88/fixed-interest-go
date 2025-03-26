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

func RolesGET(c *gin.Context) {
	var roles []models.Role
	initialize.DB.Find(&roles)

	c.JSON(200, gin.H{
		"Roles": roles,
	})
}

func RoleGET(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var role models.Role
	result := initialize.DB.First(&role, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Role": role,
	})
}

func RolePOST(c *gin.Context) {
	var body models.Role

	errBind := c.ShouldBind(&body)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": errBind.Error(),
		})
		return
	}

	result := initialize.DB.Create(&body)
	errCreate := result.Error
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": errCreate.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Role": body,
	})
}
