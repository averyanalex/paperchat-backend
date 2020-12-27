package handlers

import (
	"gorm.io/gorm"
	"github.com/bwmarrin/snowflake"
)

type Handlers struct {
	DB *gorm.DB
	SFNode *snowflake.Node
}
