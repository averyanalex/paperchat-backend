package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// ReadConfig will read configs and env to viper
func ReadConfig() {
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

	if _, err := os.Stat("/etc/nnm/base.env"); os.IsNotExist(err) {
		fmt.Println("WARNING: file base.env not found")
	} else {
		viper.SetConfigFile("/etc/nnm/base.env")
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
