package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func readConfig() {
	var err error

	viper.SetConfigFile("base.env")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("WARNING: file .env not found")
	} else {
		viper.SetConfigFile(".env")
		viper.SetConfigType("yaml")
		err = viper.MergeInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Override config parameters from environment variables if specified
	for _, key := range viper.AllKeys() {
		viper.BindEnv(key)
	}
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	readConfig()
	dsn := viper.GetString("db.user") + ":" + viper.GetString("db.password") + "@tcp(" + viper.GetString("db.address") + ":" + viper.GetString("db.port") + ")/" + viper.GetString("db.name") + "?" + viper.GetString("db.arguments")
	//fmt.Println(dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn, // data source name
		DefaultStringSize: 256, // default size for string fields
		// DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		// DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		// DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Guild{})
	db.AutoMigrate(&Channel{})
	db.AutoMigrate(&Message{})
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		db.Create(&Product{Code: name, Price: 100})
		c.String(http.StatusOK, "Hello %s", name)
	})

	r.GET("/send/:guild/:channel/:text/:user", func(c *gin.Context) {
		//channel_id := c.Param("channel")
		//guild_id := c.Param("guild")
		var channel Channel
		db.First(&channel, c.Param("channel"))
		var user User
		db.First(&user, c.Param("user"))
		db.Create(&Message{Content: c.Param("text"), User: user, Channel: channel})
		c.JSON(200, channel)
	})
	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	db.Create(&User{Name: "Ivan"})
	var gg Guild
	chanls := []Channel{{Name: "C", GuildID: gg.ID}}
	//chanls.append(Channel{Name: "C", GuildID: gg.ID})
	gg = Guild{Name: "Guild", Channels: chanls}
	db.Create(&gg)
	//db.Create(&Channel{Name: "C", GuildID: gg.ID})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	db.Delete(&product, 1)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
