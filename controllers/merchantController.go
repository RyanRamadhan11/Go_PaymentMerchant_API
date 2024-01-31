package controllers

import (
	"net/http"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"

	"github.com/RyanRamadhan11/Go_PaymentMerchant_API/initializers"
	"github.com/RyanRamadhan11/Go_PaymentMerchant_API/models"
)

// Index untuk mendapatkan semua merchant
func Index(c *gin.Context) {
	var merchants []models.Merchant
	initializers.DB.Find(&merchants)
	c.JSON(http.StatusOK, gin.H{"merchants": merchants})
}

// Show untuk mendapatkan satu merchant berdasarkan ID
func Show(c *gin.Context) {
	var merchant models.Merchant
	id := c.Param("id")

	if err := initializers.DB.First(&merchant, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"merchant": merchant})
}

// Create untuk membuat merchant baru
func Create(c *gin.Context) {
	var merchant models.Merchant

	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	initializers.DB.Create(&merchant)
	c.JSON(http.StatusOK, gin.H{"merchant": merchant})
}

// Update untuk memperbarui merchant berdasarkan ID
func Update(c *gin.Context) {
	var existingMerchant models.Merchant
	id := c.Param("id")

	if err := initializers.DB.First(&existingMerchant, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Merchant tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&existingMerchant); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := initializers.DB.Save(&existingMerchant).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Gagal menyimpan perubahan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"merchant": existingMerchant})
}

// Delete untuk menghapus merchant berdasarkan ID
func Delete(c *gin.Context) {
	var merchant models.Merchant

	var input struct {
		ID int64 `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if initializers.DB.Delete(&merchant, input.ID).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus merchant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
