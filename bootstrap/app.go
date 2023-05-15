package bootstrap

import (
	"context"
	"go-gin/app/provider"
	"go-gin/config"
	"go-gin/core/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct{}

func (*App) Run () {

	// Example Initialize the log
	// 初始化日志
	log.Init()

	// Initialize gin
	// 初始化 gin
	app := gin.New()

	// Use the default recovery
	// 使用默认的 recovery
	app.Use(gin.Recovery())

	// Configuration static file
	// 配置静态文件
	app.Static("/storage/upload", "storage/upload")

	// Load routing service
	// 加载路由服务
	(&provider.Provider{}).RouteServer(app)

	// Start http service
	// 启动 http 服务
	server := &http.Server {
		Addr: config.App["host"].(string),
		Handler: app,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	go func() {
		// Start a goroutine startup service
		// 开启一个 goroutine 启动服务
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("listen: %s\n", err)
		}
	}()

	// Wait for the interrupt signal to gracefully shut down the server,
	// setting a 5-second timeout for shutting down the server
	// 等待中断信号来优雅地关闭服务器，
	// 为关闭服务器操作设置一个5秒的超时
	// Create a channel to receive the signal
	// 创建一个接收信号的通道
	quit := make(chan os.Signal, 1)
	// kill sends the syscall.SIGTERM signal by default
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 sends the syscall.SIGINT signal. Ctrl+C is used to trigger the SIGINT signal
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 Ctrl+C 就是触发系统 SIGINT 信号
	// kill -9 sends the syscall.SIGKILL signal, but it cannot be captured, so you do not need to add it
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// Signal. Notify forwards the received syscall.SIGINT or syscall.SIGTERM signal to quit
	// signal.Notify 把收到的 syscall.SIGINT 或 syscall.SIGTERM 信号转发给 quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// The block is here and will not proceed until the above two signals are received
	// 阻塞在此，当接收到上述两种信号时才会往下执行
	<-quit
	// Create a context with a 5-second timeout
	// 创建一个5秒超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	// Gracefully shut down the service within 5 seconds
	// (processing outstanding requests before shutting down the service)
	// and timeout exits after 5 seconds
	// 5秒内优雅关闭服务
	//（将未处理完的请求处理完再关闭服务
	// 超过5秒就超时退出
	if err := server.Shutdown(ctx); err != nil {
		log.Error("Server Shutdown: ", err)
	}

}
