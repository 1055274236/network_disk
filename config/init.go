package config

import (
	"log"

	"github.com/spf13/viper"
)

var GlobalConfig ConfigStruct

type ConfigStruct struct {
	Gin struct {
		Serve struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		} `yaml:"serve"`
		Login struct {
			Ext int64 `yaml:"ext"`
		} `yaml:"login"`
	} `yaml:"gin"`
	Databases struct {
		Mysql struct {
			Account      string `yaml:"account"`
			Password     string `yaml:"password"`
			URL          string `yaml:"url"`
			Port         string `yaml:"port"`
			DbName       string `yaml:"dbName"`
			Charset      string `yaml:"charset"`
			MaxIdleConns int    `yaml:"maxIdleConns"`
			MaxOpenConns int    `yaml:"MaxOpenConns"`
		} `yaml:"mysql"`
	} `yaml:"databases"`
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
