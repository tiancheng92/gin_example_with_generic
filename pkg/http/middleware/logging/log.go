package logging

import (
	"github.com/tiancheng92/gf"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	AccessLog gin.LoggerConfig // gin的访问日志
)

func init() {
	AccessLog = gin.LoggerConfig{
		Output:    os.Stdout,
		SkipPaths: []string{"/metrics", "/healthz"}, // 可排除指定uri的access_log
		Formatter: func(params gin.LogFormatterParams) string {
			return gf.StringJoin(
				params.ClientIP,
				" - [",
				params.TimeStamp.Format("2006-01-02 15:04:05.000"),
				"] \"",
				params.Method,
				" ",
				params.Path,
				" ",
				params.Request.Proto,
				" ",
				strconv.Itoa(params.StatusCode),
				" ",
				params.Latency.String(),
				" \"",
				params.Request.UserAgent(),
				"\" ",
				params.ErrorMessage,
				"\"\n",
			)
		},
	}
	AccessLog.Output = os.Stdout
}
