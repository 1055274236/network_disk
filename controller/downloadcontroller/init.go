package downloadcontroller

import (
	"NetworkDisk/service/downloadservice"

	"github.com/gin-gonic/gin"
)

func SetupRouterOfNotRecord(engine *gin.Engine) {
	engine.GET("/temp/file/:id/:name", downloadservice.TempDownload)

	engine.GET("/static/file/:id/:name", downloadservice.DownloadFileByIndex)
}
