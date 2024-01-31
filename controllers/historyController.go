package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RyanRamadhan11/Go_PaymentMerchant_API/initializers"
	"github.com/RyanRamadhan11/Go_PaymentMerchant_API/models"
)

func GetHistories(c *gin.Context) {
	var histories []models.History

	if err := initializers.DB.Preload("User").Preload("Merchant").Find(&histories).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error fetching histories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"histories": histories})
}
