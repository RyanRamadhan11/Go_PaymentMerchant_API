package models

type Merchant struct {
    ID          int64  `gorm:"primaryKey" json:"id"`
    NamaMerchant string `gorm:"type:varchar(225)" json:"nama_merchant"`
    Deskripsi    string `gorm:"type:text" json:"deskripsi"`
    Balance     float64  // Ini adalah field balance yang menyimpan jumlah uang pada akun merchant
}

