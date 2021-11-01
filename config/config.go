package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Setup(env, configPath string) {
	if env == "" {
		panic(fmt.Errorf("📌📌📌📌📌 PLEASE SET `ENV` ex. `export ENV=dev` 📌📌📌📌📌"))
	}
	viper.SetConfigType("json")
	viper.SetConfigName("env." + env)
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Debug()
}
