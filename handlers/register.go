package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/paper-chat/nnm/models"
	"github.com/paper-chat/nnm/utils"
)

// Register create specific accaunt
func (h Handlers) Register(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(400, &models.Result{Error: "Please, provide name", Code: "reg_no_name"})
		return
	}
	email := c.PostForm("email")
	if email == "" {
		c.JSON(400, &models.Result{Error: "Please, provide email", Code: "reg_no_email"})
		return
	}
	rawPassword := c.PostForm("password")
	if name == "" {
		c.JSON(400, &models.Result{Error: "Please, provide password", Code: "reg_no_password"})
		return
	}
	hashedPassword := utils.HashPassword(rawPassword)
	msg := &models.User{ID: uint(h.SFNode.Generate().Int64()), Name: name, Email: email, Password: hashedPassword}
	h.DB.Create(&msg)
	c.JSON(200, models.Result{Code: "OK"})
}