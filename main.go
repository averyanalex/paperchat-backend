package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fx.New(
		fx.Provide(
			newLogger,
			newGin,
			newDB,
		),
		fx.Invoke(register),
	)
	app.Run()
}

func newGin(lifecycle fx.Lifecycle, logger *log.Logger) *gin.Engine {
	logger.Print("Executing NewGin")
	router := gin.Default()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return server.Shutdown(ctx)
			},
		},
	)
	return router
}

func newLogger() *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")
	return logger
}

func register(router *gin.Engine, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello!")
	})
}

func newDB(logger *log.Logger) *gorm.DB {
	logger.Print("Executing newDB")
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
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Guild{})
	db.AutoMigrate(&Channel{})
	db.AutoMigrate(&Message{})
	return db
}

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
