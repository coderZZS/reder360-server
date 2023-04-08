package routes

import (
	"net/http"
	"reader360Server/src/controller"
	"reader360Server/src/logger"
	"reader360Server/src/middleware"

	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	_ "reader360Server/docs"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()
	// 挂载文档
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// 日志中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(context *gin.Context) {
		// context.String(http.StatusOK, "ok")
		slice := []interface{}{}
		context.JSON(http.StatusOK, slice)
	})
	r.GET("/ping", middleware.JWTAuthMiddleware(), func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})

	// user模块
	{
		user := r.Group("/api/user")
		user.POST("/signUp", controller.UserSignUp)
		user.POST("/login", controller.UserLogin)
	}
	r.Use(middleware.JWTAuthMiddleware()) // 设置鉴权中间件
	return r
}
