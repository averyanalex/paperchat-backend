package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Attachment contain data
type Attachment struct {
	gorm.Model
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}
