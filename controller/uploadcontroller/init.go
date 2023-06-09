package uploadcontroller

import (
	"NetworkDisk/service/uploadservice"

	"github.com/gin-gonic/gin"
)

func SetupRoute(engine *gin.Engine, userVerified *gin.RouterGroup) {
	userVerified.POST("/upload/base", uploadservice.BaseUpload)

	userVerified.POST("/upload/mkdir", uploadservice.Mkdir)
}
