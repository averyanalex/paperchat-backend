package models

import (
	"gorm.io/gorm"
)

// Guild contain guild data
type Guild struct {
	gorm.Model
	ID       uint      `json:"id" gorm:"primary_key"`
	Name     string    `json:"name" gorm:"size:30"`
	Users    []*User   `gorm:"many2many:user_guilds"`
	Channels []Channel `gorm:"many2many:guild_channels"`
}
