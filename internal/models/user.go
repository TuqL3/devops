package models

import (
	"gorm.io/gorm"
)

// User model cho người dùng
type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"-"`
	Role     string `json:"role"`
	Active   bool   `json:"active" gorm:"default:true"`
}
