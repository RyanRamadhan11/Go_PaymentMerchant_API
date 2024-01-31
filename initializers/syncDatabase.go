// initializers/migrations.go
package initializers

import (
	"github.com/RyanRamadhan11/Go_PaymentMerchant_API/models"
)

// SyncDatabase untuk melakukan migrasi database
func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Merchant{})
	DB.AutoMigrate(&models.History{})
}
