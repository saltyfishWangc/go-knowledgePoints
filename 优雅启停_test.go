package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

// Test优雅启停
// 测试方法：将程序编译在终端运行，发送/请求，然后迅速的执行kill命令关闭服务，会发现5秒后还是收到了前面请求的响应：Welcome Gin Server
func Test优雅启停(t *testing.T) {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.String(http.StatusOK, "Welcome Gin Server")
	})

	src := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := src.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Logf("listen: %s\n", err)
		}
	}()

	// 定义一个通道来接收中断信号的系统调用
	quit := make(chan os.Signal, 1)

	// kill 默认会发送syscall.SIGTERM信号
	// kill -2 发送syscall.SIGINT信号，我们常用的Ctrl + C就是触发系统SIGINT信号
	// kill -9 发送syscall.SIGKILL信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的syscall.SIGINT或syscall.SIGTERM信号转发给quit通道
	// 注意：signal.Notify()并不会阻塞
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 主协程阻塞在这里，当收到上述两种信号时才会往下执行
	<-quit
	t.Log("Shutdown Server...")

	// 创建一个5秒超时的context
	context, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	// 5秒内优雅关闭服务(将未处理完的请求处理完再关闭服务)，超过5秒就超时退出
	if err := src.Shutdown(context); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}
