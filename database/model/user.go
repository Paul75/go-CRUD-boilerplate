package model

import (
	"github.com/jinzhu/gorm"
)

// User model - `users` table
type User struct {
	gorm.Model
	Name  string `json:"Name"`
	Email string `json:"Email"`
	Posts []Post `gorm:"many2many:user_posts";"foreignkey:UserID";"association_foreignkey:ID"`
}
