package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handlers struct {
	DB *gorm.DB
}

func SetupHandlers(db *gorm.DB) *Handlers {
	return &Handlers{DB: db}
}

func (h Handlers) Ping(c *gin.Context) {
	c.String(200, "Pong!")
}
