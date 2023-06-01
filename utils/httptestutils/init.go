package httptestutils

import (
	"NetworkDisk/config"
	"NetworkDisk/controller"

	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func init() {
	gin.SetMode(gin.DebugMode)
	Engine = gin.New()
	controller.SetupRouter(Engine)
	go Engine.Run(config.GlobalConfig.Gin.Serve.Host + config.GlobalConfig.Gin.Serve.Port)
}
