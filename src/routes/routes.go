package routes

import (
	"gin-cli/src/controller"
	"gin-cli/src/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	r.POST("/signUp", controller.UserSignUp)
	return r
}
