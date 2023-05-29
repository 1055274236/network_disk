package config

import (
	"log"

	"github.com/spf13/viper"
)

var GlobalConfig ConfigStruce

type ConfigStruce struct {
	Gin struct {
		Serve struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		} `yaml:"serve"`
	} `yaml:"gin"`
}

func init() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config file failed, %v", err)
	}
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Printf("unmarshal config file failed, %v", err)
	}
}
