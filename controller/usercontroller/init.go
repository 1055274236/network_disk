package usercontroller

import (
	"NetworkDisk/service/userservice"

	"github.com/gin-gonic/gin"
)

func SetupRoute(engine *gin.Engine) {
	engine.POST("/login", userservice.Login)
	engine.POST("/signin", userservice.SignIn)
}
