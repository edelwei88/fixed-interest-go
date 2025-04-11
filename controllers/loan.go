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

func LoansGET(c *gin.Context) {
	var loans []models.Loan
	result := initialize.DB.Preload("LoanPayments").Preload("LoanType").Find(&loans)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"Loans": loans,
	})
}

func LoanGET(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var loan models.Loan
	result := initialize.DB.First(&loan, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
		return
	}
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Loan": loan,
	})
}

func LoanPOST(c *gin.Context) {
	var body models.Loan
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
		"Loan": body,
	})
}

func LoanPATCH(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var body models.Loan
	err = c.ShouldBind(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var loan models.Loan
	result := initialize.DB.First(&loan, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
		return
	}
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = initialize.DB.Model(&loan).Omit("ID").Updates(&body)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Loan": loan,
	})
}

func LoanDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result := initialize.DB.Delete(&models.Loan{}, id)
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
