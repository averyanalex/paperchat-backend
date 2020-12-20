package main

import (
	"time"

	"gorm.io/gorm"
)

// User is a structure than contain user data
type User struct {
	gorm.Model
	ID      uint     `json:"id" gorm:"primary_key"`
	Name    string   `json:"name" gorm:"size:30"`
	Guilds  []*Guild `gorm:"many2many:user_guilds"`
	Friends []*User  `gorm:"many2many:user_friends"`
}

// Guild contain guild data
type Guild struct {
	gorm.Model
	ID       uint      `json:"id" gorm:"primary_key"`
	Name     string    `json:"name" gorm:"size:30"`
	Users    []*User   `gorm:"many2many:user_guilds"`
	Channels []Channel `gorm:"many2many:guild_channels"`
}

// Channel containt channel info
type Channel struct {
	gorm.Model
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name" gorm:"size:30"`
	GuildID uint
}

// Message info
type Message struct {
	gorm.Model
	ID      uint   `json:"id" gorm:"primary_key"`
	Content string `json:"content"`
	//UserID    uint
	//User      User
	//ChannelID uint
	//Channel   Channel
}

// MessageToUser structure sent to client when he request messages
type MessageToUser struct {
	ID          uint        `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time   `json:"time"`
	Content     string      `json:"content"`
}

// ClientError is an error structure sent to client if he wrong
type ClientError struct {
	Error string `json:"error"`
}
