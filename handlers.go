package main

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"crypto/md5"
	"fmt"
)

// Handlers contain DB and have methods to process requests
type Handlers struct {
	DB *gorm.DB
}

// SetupHandlers create new Handlers with given DB
func SetupHandlers(db *gorm.DB) *Handlers {
	return &Handlers{DB: db}
}

// Ping fuctions will help to test API status
func (h Handlers) Ping(c *gin.Context) {
	c.JSON(200, &Result{Message: "Pong!"})
}

// Send will save sent message
func (h Handlers) Send(c *gin.Context) {
	content := c.Request.FormValue("message")
	if content != "" {
		file, error := c.FormFile("file")
		if error != nil {
			fmt.Println(error)
			// filename := filepath.Base(file.Filename)
			// Upload the file to specific dst.
			uploadUUID := uuid.New()
			err := os.Mkdir("upload/"+uploadUUID.String(), 0755)
			fmt.Println(err)
			c.SaveUploadedFile(file, "upload/"+uploadUUID.String()+"/"+filepath.Base(file.Filename))
			h.DB.Create(&Attachment{UUID: uploadUUID, Name: filepath.Base(file.Filename)})
		}
		msg := &Message{Content: content, IP: c.ClientIP()}
		h.DB.Create(msg)
		c.Status(200)
	} else {
		c.JSON(400, &Result{Error: "Empty Message"})
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
			c.JSON(400, &Result{Error: "Failed to parse count to int"})
			return
		}
		if count > 75 {
			c.JSON(400, &Result{Error: "Messages count to large. Maximum is 75"})
			return
		}
	}
	startString, startStated := c.GetQuery("start")
	if startStated {
		var start int
		start, err = strconv.Atoi(startString)
		if err != nil {
			c.JSON(400, &Result{Error: "Failed to parse start to int"})
			return
		}
		h.DB.Model(&Message{}).Where("id <= ?", start).Order("id DESC").Limit(count).Find(&msg)
	} else {
		h.DB.Model(&Message{}).Order("id DESC").Limit(count).Find(&msg)
	}
	c.JSON(200, msg)
}

// Register handle user registration
func (h Handlers) Register(c *gin.Context) {
	h.DB.Create(&User{Name: c.PostForm("name"), Password: c.PostForm("password"), Email: c.PostForm("mail")})
	c.String(200, "OK")
}

func (h Handlers) Upload(c *gin.Context) {
	// single file
	file, error := c.FormFile("file")
	fmt.Println(error)
	// filename := filepath.Base(file.Filename)
	// Upload the file to specific dst.
	uploadUUID := uuid.New()
	err := os.Mkdir("upload/"+uploadUUID.String(), 0755)
	fmt.Println(err)
	c.SaveUploadedFile(file, "upload/"+uploadUUID.String()+"/"+filepath.Base(file.Filename))
	h.DB.Create(&Attachment{UUID: uploadUUID, Name: filepath.Base(file.Filename)})
	c.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

// Sabotage tochno ne sabotiruet vse
func (h Handlers) Sabotage(c *gin.Context) {
	pass, passGiven := c.GetQuery("password")
	if passGiven {

		h := md5.New()
		h.Write([]byte(pass))
		bs := h.Sum(nil)
		if fmt.Sprintf("%x", bs) == "7a4692db8b97a98afbece98e08014205" {
			c.Status(200)
			panic("импостор амонг ас")
		}

	}
	c.JSON(418, &Result{Error: "Ты совсем чайник что-ли? Пароль нормальный введи"})
}
