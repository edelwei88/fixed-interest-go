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

func LoanPaymentsGET(c *gin.Context) {
	var lps []models.LoanPayment
	result := initialize.DB.Find(&lps)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"LoanPayments": lps,
	})
}

func LoanPaymentGET(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var lp models.LoanPayment
	result := initialize.DB.First(&lp, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
		return
	}
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"LoanPayment": lp,
	})
}

func LoanPaymentPOST(c *gin.Context) {
	var body models.LoanPayment
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
		"LoanPayment": body,
	})
}

func LoanPaymentPATCH(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var body models.LoanPayment
	err = c.ShouldBind(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var lp models.LoanPayment
	result := initialize.DB.First(&lp, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
		return
	}
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = initialize.DB.Model(&lp).Omit("ID", "LoanID").Updates(&body)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"LoanPayment": lp,
	})
}

func LoanPaymentDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result := initialize.DB.Delete(&models.LoanPayment{}, id)
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
