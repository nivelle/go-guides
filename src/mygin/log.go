package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	//禁用控制台颜色，当你将日志写入到文件的时候,不需要控制台颜色
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	//同时写入日志文件和控制台上显示
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
	router.Run(":8000")
}
