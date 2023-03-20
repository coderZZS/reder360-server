package routes

import (
	"fmt"
	"gin-cli/src/logger"
	"gin-cli/src/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(context *gin.Context) {
		fmt.Println(settings.Conf.Version, "------------------")
		context.String(http.StatusOK, settings.Conf.Version)
	})
	return r
}
