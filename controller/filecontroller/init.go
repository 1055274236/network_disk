package filecontroller

import (
	"NetworkDisk/service/fileservice"

	"github.com/gin-gonic/gin"
)

func SetupRoute(engine *gin.Engine, userVerified *gin.RouterGroup) {
	userVerified.POST("/getfileindex", fileservice.GetFile)
}
