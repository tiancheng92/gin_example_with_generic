package main

import (
	"fmt"
	"gin_example_with_generic/config"
	"gin_example_with_generic/pkg/log"
	"gin_example_with_generic/pkg/validator"
	"gin_example_with_generic/server"
	"gin_example_with_generic/store"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"
	"math/rand"
	"os"
	"time"
)

//go:generate codegen -type=int ../pkg/ecode

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:     "gin_example_with_generic",
		Version: "v1.0.0",
		Short:   "gin example",
		Long:    "gin example with generic",
		Run: func(cmd *cobra.Command, _ []string) {
			config.Init()
			// 初始化appLog
			log.Init()
			// 参数校验国际化
			validator.Init()
			// 初始化数据库
			store.Init()
			// 启动Web服务
			server.Run()
		},
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
