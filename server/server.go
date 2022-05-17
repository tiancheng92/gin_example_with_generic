package server

//go:generate swag fmt -g server.go

import (
	"context"
	"gin_example_with_generic/config"
	"gin_example_with_generic/pkg/log"
	"gin_example_with_generic/router"

	"github.com/tiancheng92/gf"

	"net/http"
	"os"
	"os/signal"
	"strconv"

	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	conf := config.GetConf()
	switch conf.Server.Mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	s := &http.Server{
		Addr:    gf.StringJoin(conf.Server.ServiceHost, ":", strconv.Itoa(conf.Server.ServicePort)), // 监听地址
		Handler: router.InitRouter(),                                                                // 处理器
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Warnf("Listen Stop, reason: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGSEGV,
		syscall.SIGABRT,
		syscall.SIGILL,
		syscall.SIGFPE)
	<-quit

	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %v", err.Error())
	}

	log.Info("Server exiting")
	close(quit)
}
