package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"net/http"
	"os"
	"io"
)

func main() {

	router := gin.New()
	f, _ := os.Create("gin.log")
	//同时写入日志文件和控制台上显示
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//记录日志的颜色
	gin.ForceConsoleColor()
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage, )
	}))
	router.Use(gin.Recovery())

	router.GET("ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	router.Run(":9000")

}
