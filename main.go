package main

import (
	"NetworkDisk/config"
	"NetworkDisk/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	controller.SetupRouter(engine)
	engine.Run(config.GlobalConfig.Gin.Serve.Host + config.GlobalConfig.Gin.Serve.Port)
}
