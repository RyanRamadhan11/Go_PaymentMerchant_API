// history.go
package models

type History struct {
	ID         int64   `json:"id"`
	UserID int64   `json:"user_id"`
	MerchantID int64   `json:"merchant_id"`
	Amount     float64 `json:"amount"`

	User   User `json:"user" gorm:"foreignKey:UserID"`
	Merchant   Merchant `json:"merchant" gorm:"foreignKey:MerchantID"`
}
