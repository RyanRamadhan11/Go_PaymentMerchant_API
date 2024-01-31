package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/RyanRamadhan11/Go_PaymentMerchant_API/initializers"
	"github.com/RyanRamadhan11/Go_PaymentMerchant_API/models"
)

// Payment untuk melakukan pembayaran
func Payment(c *gin.Context) {
	var paymentData struct {
		UserID     int64   `json:"user_id"`
		MerchantID int64   `json:"merchant_id"`
		Amount     float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&paymentData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request format"})
		return
	}

	// Mencari user yang sesuai dengan ID
	var user models.User
	if err := initializers.DB.First(&user, paymentData.UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User not found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error querying user data"})
		return
	}

	// Mencari merchant yang sesuai dengan ID
	var merchant models.Merchant
	if err := initializers.DB.First(&merchant, paymentData.MerchantID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Merchant not found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error querying merchant data"})
		return
	}

	// Melakukan pembayaran
	if user.Balance >= paymentData.Amount {
		// Mengurangi saldo user
		user.Balance -= paymentData.Amount

		// Menambah saldo merchant
		merchant.Balance += paymentData.Amount

		// Menyimpan perubahan data user
		if err := initializers.DB.Save(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error updating user data"})
			return
		}

		// Menyimpan perubahan data merchant
		if err := initializers.DB.Save(&merchant).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error updating merchant data"})
			return
		}

		// Menyimpan history pembayaran
		var history models.History
		history.UserID = paymentData.UserID
		history.MerchantID = paymentData.MerchantID
		history.Amount = paymentData.Amount

		if err := initializers.DB.Create(&history).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error creating payment history"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Payment success"})
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Insufficient balance"})
	}
}
