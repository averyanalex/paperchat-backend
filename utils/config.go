package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// ReadConfig will read configs and env to viper
func ReadConfig() {
	// var err error

	// viper.SetConfigFile("base.env")
	// viper.SetConfigType("yaml")
	// err = viper.ReadInConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	if _, err := os.Stat("/etc/paper-chat/api.yml"); os.IsNotExist(err) {
		fmt.Println("Config file /etc/paper-chat/api.yml not found")
	} else {
		viper.SetConfigFile("/etc/paper-chat/api.yml")
		viper.SetConfigType("yaml")
		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if _, err := os.Stat("~/.config/paper-chat/api.yml"); os.IsNotExist(err) {
		fmt.Println("Config file ~/.config/paper-chat/api.yml not found")
	} else {
		viper.SetConfigFile("~/.config/paper-chat/api.yml")
		viper.SetConfigType("yaml")
		err = viper.MergeInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if _, err := os.Stat("api.yml"); os.IsNotExist(err) {
		fmt.Println("Config file api.yml not found")
	} else {
		viper.SetConfigFile("api.yml")
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
