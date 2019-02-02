package config

import (
	"fmt"

	. "github.com/spf13/viper"
)

func InitViper() {
	SetConfigName("config")         // name of config file (without extension)
	AddConfigPath("/etc/appname/")  // path to look for the config file in
	AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	AddConfigPath(".")              // optionally look for config in the working directory
	err := ReadInConfig()           // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
