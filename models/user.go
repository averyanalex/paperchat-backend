package models

import (
	"gorm.io/gorm"
)

// User is a structure than contain user data
type User struct {
	gorm.Model
	ID       uint     `json:"id" gorm:"primary_key"`
	Name     string   `json:"name" gorm:"size:30"`
	Guilds   []*Guild `gorm:"many2many:user_guilds"`
	Friends  []*User  `gorm:"many2many:user_friends"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
}
