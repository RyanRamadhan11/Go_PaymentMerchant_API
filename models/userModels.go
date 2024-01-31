package models

import "gorm.io/gorm"

type User struct {
    gorm.Model    
    Name     string
    Email    string `gorm:"unique"`    
    Password string
    Address  string
    Balance  float64  // Ini adalah field balance yang menyimpan jumlah uang pada akun pengguna
}
