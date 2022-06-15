package main

import (
	"fmt"
	"gin_example_with_generic/config"
	"gin_example_with_generic/pkg/log"
	"gin_example_with_generic/pkg/validator"
	"gin_example_with_generic/server"
	"gin_example_with_generic/store"
	"github.com/spf13/cobra"
	"github.com/tiancheng92/gf"
	_ "go.uber.org/automaxprocs"
	"math/rand"
	"os"
	"time"
)

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
			// 初始化数据库
			store.Init()
			// 启动Web服务
			server.Run()
		},
	}

	defaultConfigPath := "./config_file/local.yaml"
	env := os.Getenv("APP_ENV") // 环境变量
	if env != "" {
		defaultConfigPath = gf.StringJoin("./config_file/", env, ".yaml")
	}

	rootCmd.PersistentFlags().StringVarP(&configPath, "config_path", "f", defaultConfigPath, "config of the Gin Example")
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
