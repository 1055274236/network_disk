package controller

import (
	"NetworkDisk/controller/downloadcontroller"
	"NetworkDisk/controller/uploadcontroller"
	"NetworkDisk/controller/usercontroller"
	"NetworkDisk/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {
	downloadcontroller.SetupRouterOfNotRecord(engine)
	engine.Use(middleware.Recovery(), middleware.Logger(), middleware.Cors())
	userVerified := engine.Group("/", middleware.UserVerify())
	usercontroller.SetupRoute(engine)
	uploadcontroller.SetupRoute(engine, userVerified)
}
