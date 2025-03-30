package middlewares

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/edelwei88/fixed-interest-go/initialize"
	"github.com/edelwei88/fixed-interest-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BearerTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		reqToken := strings.Split(bearerToken, " ")[1]

		var existingToken models.Token

		result := initialize.DB.Where(&models.Token{Token: reqToken}).First(&existingToken)
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

		c.Next()
	}
}
