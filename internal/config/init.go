package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %v", err.Error())
	}
}
