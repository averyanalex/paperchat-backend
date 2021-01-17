package utils

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// ReadConfig will read configs and env to viper
func ReadConfig() {
	viper.SetConfigName("api.yml")                  // name of config file (without extension)
	viper.SetConfigType("yaml")                     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/paper-chat/")         // path to look for the config file in
	viper.AddConfigPath("$HOME/.config/paper-chat") // call multiple times to add many search paths
	viper.AddConfigPath(".")                        // optionally look for config in the working directory
	err := viper.ReadInConfig()                     // Find and read the config file
	if err != nil {                                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("paper")
	viper.AutomaticEnv()
}
