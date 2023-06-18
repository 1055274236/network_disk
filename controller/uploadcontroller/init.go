package uploadcontroller

import (
	"NetworkDisk/service/uploadservice"

	"github.com/gin-gonic/gin"
)

func SetupRoute(engine *gin.Engine, userVerified *gin.RouterGroup) {
	userVerified.POST("/upload/base", uploadservice.BaseUpload)
	userVerified.POST("/upload/file/:id", uploadservice.UploadById)

	userVerified.POST("/upload/mkdir", uploadservice.Mkdir)
	userVerified.POST("/upload/createfile", uploadservice.CreateFileIndex)
	engine.GET("/status/file", uploadservice.GetFileStatus)
	userVerified.GET("/status/filebyid", uploadservice.GetFileStatusById)
}
