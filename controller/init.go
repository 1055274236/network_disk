package controller

import (
	"NetworkDisk/controller/usercontroller"
	"NetworkDisk/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {
	engine.Use(middleware.Recovery(), middleware.Logger(), middleware.Cors())
	// userVerified := engine.Group("/", middleware.UserVerify())
	usercontroller.SetupRoute(engine)
}
