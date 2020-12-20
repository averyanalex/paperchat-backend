package main

import (
	"strconv"

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

// Send will save sent message
func (h Handlers) Send(c *gin.Context) {
	content, contentGiven:= c.GetQuery("content")
	if contentGiven {
		msg := &Message{Content: content}
		h.DB.Create(msg)
		c.Status(200)
	} else {
		c.JSON(400, &ClientError{Error: "Empty Message"})
	}
}

// GetMsgs will send requested messages to client
func (h Handlers) GetMsgs(c *gin.Context) {
	var msg []MessageToUser
	countString, countStated := c.GetQuery("count")
	count := 25
	var err error
	if countStated {
		count, err = strconv.Atoi(countString)
		if err != nil {
			c.JSON(400, &ClientError{Error: "Failed to parse count to int"})
			return
		}
		if count > 75 {
			c.JSON(400, &ClientError{Error: "Messages count to large. Maximum is 75"})
			return
		}
	}
	startString, startStated := c.GetQuery("start")
	if startStated {
		var start int
		start, err = strconv.Atoi(startString)
		if err != nil {
			c.JSON(400, &ClientError{Error: "Failed to parse start to int"})
			return
		}
		h.DB.Model(&Message{}).Where("id <= ?", start).Order("id DESC").Limit(count).Find(&msg)
	} else {
		h.DB.Model(&Message{}).Order("id DESC").Limit(count).Find(&msg)
	}
	c.JSON(200, msg)
}
