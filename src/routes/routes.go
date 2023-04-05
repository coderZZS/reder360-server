package routes

import (
	"gin-cli/src/controller"
	"gin-cli/src/logger"
	"gin-cli/src/middleware"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	_ "gin-cli/docs"

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
	// community
	community := r.Group("/api/community")
	{
		community.GET("/", controller.CommunityTagList)
		community.GET("/:id", controller.CommunityDetail)
	}

	// post
	post := r.Group("/api/post")
	{
		post.POST("/add", controller.PostAdd)
		post.GET("/:id", controller.PostDetail)
		post.GET("/list", controller.PostList)
		post.GET("/list/type", controller.PostListByType)
	}

	// vote
	vote := r.Group("/api/vote")
	{
		vote.POST("/", controller.VotePost)
	}
	return r
}
