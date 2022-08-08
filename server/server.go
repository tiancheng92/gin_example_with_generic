package server

import (
	"context"
	"gin_example_with_generic/config"
	"gin_example_with_generic/pkg/log"
	"gin_example_with_generic/router"
	"github.com/gin-gonic/gin"
	"github.com/tiancheng92/gf"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func runServer(stop, ready chan struct{}) {
	conf := config.GetConf()
	switch conf.Server.Mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	// 定义http服务
	s := &http.Server{
		Addr:    gf.StringJoin(conf.Server.ServiceHost, ":", strconv.Itoa(conf.Server.ServicePort)), // 监听地址
		Handler: router.InitRouter(),                                                                // 处理器
	}

	// 开始监听
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Warnf("Listen Stop, reason: %s", err)
		}
	}()

	<-stop

	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %v", err.Error())
	}
	log.Info("Server exiting")
	ready <- struct{}{}
}

func Run() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGSEGV,
		syscall.SIGABRT,
		syscall.SIGILL,
		syscall.SIGFPE,
	)

	stop := make(chan struct{})
	ready := make(chan struct{})
	go runServer(stop, ready)
	go func() {
		for {
			select {
			case <-config.HotUpdateForServer:
				stop <- struct{}{}
				<-ready
				go runServer(stop, ready)
				log.Info("Gin Server 热更新完成。")
			}
		}
	}()
	<-quit
}
