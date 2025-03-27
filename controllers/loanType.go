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

func LoanTypesGET(c *gin.Context) {
	var lts []models.LoanType
	result := initialize.DB.Find(&lts)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"LoanTypes": lts,
	})
}

func LoanTypeGET(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var lt models.LoanType
	result := initialize.DB.First(&lt, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
		return
	}
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"LoanType": lt,
	})
}

func LoanTypePOST(c *gin.Context) {
	var body models.LoanType
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
		"LoanType": body,
	})
}

func LoanTypePATCH(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var body models.LoanType
	err = c.ShouldBind(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var lt models.LoanType
	result := initialize.DB.First(&lt, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
		return
	}
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = initialize.DB.Model(&lt).Omit("ID").Updates(&body)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"LoanType": lt,
	})
}

func LoanTypeDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result := initialize.DB.Delete(&models.LoanType{}, id)
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
