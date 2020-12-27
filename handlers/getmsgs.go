package handlers

import (
	"strconv"
	"github.com/paper-chat/nnm/models"
	"github.com/gin-gonic/gin"
)

// GetMsgs will send requested messages to client
func (h Handlers) GetMsgs(c *gin.Context) {
	var msg []models.MessageToUser
	countString, countStated := c.GetQuery("count")
	count := 25
	var err error
	if countStated {
		count, err = strconv.Atoi(countString)
		if err != nil {
			c.JSON(400, &models.Result{Error: "Failed to parse count to int"})
			return
		}
		if count > 75 {
			c.JSON(400, &models.Result{Error: "Messages count to large. Maximum is 75"})
			return
		}
	}
	startString, startStated := c.GetQuery("start")
	if startStated {
		var start int
		start, err = strconv.Atoi(startString)
		if err != nil {
			c.JSON(400, &models.Result{Error: "Failed to parse start to int"})
			return
		}
		h.DB.Model(&models.Message{}).Where("id <= ?", start).Order("id DESC").Limit(count).Find(&msg)
	} else {
		h.DB.Model(&models.Message{}).Order("id DESC").Limit(count).Find(&msg)
	}
	c.JSON(200, msg)
}
