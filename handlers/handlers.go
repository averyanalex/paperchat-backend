package handlers

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handlers struct {
	DB     *gorm.DB
	SFNode *snowflake.Node
}

// Setup will add handlers to api
func Setup(router *gin.Engine, handlers *Handlers) {
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello!")
	})
	// router.GET("/ping", handlers.Ping)
	router.POST("/send/:msg", handlers.Send)
	router.GET("/get", handlers.GetMsgs)
	router.POST("/register", handlers.Register)
	// router.POST("/reg", handlers.Register)
	// router.POST("/upload", handlers.Upload)
	// router.GET("/sabotage", handlers.Sabotage)
}
