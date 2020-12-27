package models

import (
	"time"
)

// Message info
type Message struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Content string `json:"content"`
	IP      string `json:"ip"`
	//UserID    uint
	//User      User
	//ChannelID uint
	//Channel   Channel
}

// MessageToUser structure sent to client when he request messages
type MessageToUser struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"time"`
	Content   string    `json:"content"`
	IP        string    `json:"ip"`
}
