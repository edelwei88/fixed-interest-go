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
	result := initialize.DB.Find(&roles)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

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
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Role": role,
	})
}

func RolePOST(c *gin.Context) {
	var body models.Role
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
		"Role": body,
	})
}

func RolePATCH(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var body models.Role
	err = c.ShouldBind(&body)
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
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = initialize.DB.Model(&role).Omit("ID").Updates(&body)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Role": role,
	})
}

func RoleDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result := initialize.DB.Delete(&models.Role{}, id)
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
