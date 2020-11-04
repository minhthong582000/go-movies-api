package middleware

import (
	"io"
	"time"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",param.ClientIP,param.TimeStamp.Format(time.RFC822),param.Method,param.Path,param.StatusCode,param.Latency)
	})
}