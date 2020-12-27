package models

import (
	"gorm.io/gorm"
)

// Channel containt channel info
type Channel struct {
	gorm.Model
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name" gorm:"size:30"`
	GuildID uint
}
