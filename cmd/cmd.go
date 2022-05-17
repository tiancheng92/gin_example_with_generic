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

//go:generate swag fmt -g ../server/server.go

//go:generate swag init -g ../server/server.go -o ../docs --parseDependency --parseInternal --generatedTime --parseDepth 4

//go:generate codegen -type=int ../pkg/ecode

var (
	configPath string
	rootCmd    *cobra.Command
)

func init() {
	rootCmd = &cobra.Command{
		Use:     "gin_example_with_generic",
		Version: "v1.0.0",
		Short:   "gin example",
		Long:    "gin example with generic",
		Run: func(cmd *cobra.Command, args []string) {
			config.Init(configPath)
			// 初始化appLog
			log.Init(config.GetConf().Log.Level)
			// 参数校验国际化
			validator.Init(config.GetConf().I18n.Locale)
			// 初始化默认数据库
			store.InitDefaultDB()
			// 启动Web服务
			server.Run()
		},
	}
	rootCmd.PersistentFlags().StringVarP(&configPath, "config_path", "f", "./config_file/local.yaml", "config of the Gin Example")
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
