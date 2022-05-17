package crossDomain

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CrossDomain 处理跨域
func CrossDomain() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"X-Total-Count"},
	})
}
