package main

import (
	"context"
	"fmt"
	"gin-cli/src/dao/mysql"
	"gin-cli/src/dao/redis"
	"gin-cli/src/logger"
	"gin-cli/src/pkg/snowflake"
	"gin-cli/src/routes"
	"gin-cli/src/settings"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// title gin-cli接口文档
// @version 1.0.0
// description gin-cli的接口文档
// termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8088
// @BasePath
func main() {
	// 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("settings init failed, err: %v\n", err)
		return
	}
	// 2.初始化日志

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("logger init failed, err: %v\n", err)
		return
	}
	//2.1延迟同步日志
	defer zap.L().Sync()
	// 3.初始化数据库
	// 3.1 mysql
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("mysql init failed, err: %v\n", err)
		return
	}
	// 释放数据库
	defer mysql.Close()

	// 3.2 redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("redis init failed, err: %v\n", err)
		return
	}
	// 释放数据库
	defer redis.Close()

	// 4.注册路由
	r := routes.Init()
	// 5.启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 6.初始化雪花算法
	snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID)

	// 等待中断信号，关闭服务，为关闭服务设置一个5秒的超时
	quit := make(chan os.Signal, 1)
	// kill默认会发送syscall.SIGTERM 信号
	// kill -2 发送syscall.SIGINT 信号，例如我们常用的ctrl+c就是触发这个信号
	// kill -9 发送syscall.SIGKILL信号，但是不能被捕获，所以不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞在这里，等待发送销毁信号
	zap.L().Info("Shutdown Server...")
	//	创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//	5秒内优雅关闭服务（将未处理完的请求处理完毕再关闭服务）， 超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server shutdown", zap.Error(err))
	}
	zap.L().Fatal("Server exiting...")
}
