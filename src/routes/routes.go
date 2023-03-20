package routes

import (
	"fmt"
	"gin-cli/src/logger"
	"gin-cli/src/pkg/snowflake"
	"gin-cli/src/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()
	fmt.Println(settings.Conf.Version, "------------------")
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	r.GET("/test", func(context *gin.Context) {
		context.String(http.StatusOK, "request id:%d", snowflake.GenID())
	})
	return r
}
