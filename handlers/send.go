package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/paper-chat/nnm/models"
)

// Send will save sent message
func (h Handlers) Send(c *gin.Context) {
	//fmt.Println(c.GetRawData())
	content := c.DefaultPostForm("message", "")
	if content != "" {
		// file, err := c.FormFile("file")
		// if err != nil {
		// 	panic(err)
		// }
		// if true {
		// 	uploadUUID := uuid.New()
		// 	err = os.Mkdir("/tmp/paper", 0755)
		// 	// if err != nil {
		// 	// 	panic(err)
		// 	// }
		// 	err = os.Mkdir("/tmp/paper/upload", 0755)
		// 	// if err != nil {
		// 	// 	panic(err)
		// 	// }
		// 	err = os.Mkdir("/tmp/paper/upload/attachments", 0755)
		// 	// if err != nil {
		// 	// 	panic(err)
		// 	// }
		// 	err = os.Mkdir("/tmp/paper/upload/attachments/"+uploadUUID.String(), 0755)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	filePath := "attachments/" + uploadUUID.String() + "/" + filepath.Base(file.Filename)
		// 	fullFilePath := "/tmp/paper/upload/" + filePath
		// 	c.SaveUploadedFile(file, fullFilePath)
		// 	h.DB.Create(&models.Attachment{UUID: uploadUUID, Name: filepath.Base(file.Filename)})
		// 	openedFile, err := os.Open(fullFilePath)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	go utils.UploadFile(filePath, *openedFile)
		// }
		msg := &models.Message{Content: content, IP: c.ClientIP(), ID: uint(h.SFNode.Generate().Int64()), Chat: c.Param("id")}
		h.DB.Create(msg)
		c.JSON(200, &models.Result{})
	} else {
		c.JSON(400, &models.Result{Error: "Empty Message"})
	}
}
