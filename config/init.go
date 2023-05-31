package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

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
	viper.AddConfigPath(getCurrentAbPath())
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config file failed, %v", err)
	}
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Printf("unmarshal config file failed, %v", err)
	}
}

// 最终方案-全兼容
func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	if strings.Contains(dir, getTmpDir()) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
