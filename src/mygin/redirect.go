package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/test", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://baidu.com")
	})

	router.POST("test2", func(context *gin.Context) {
		context.Redirect(http.StatusFound, "/foo")
	})

	router.GET("/foo", func(context *gin.Context) {
		context.String(http.StatusOK, "redirect ok")
	})

	router.GET("/test3", func(context *gin.Context) {
		context.Request.URL.Path = "/test4"
		router.HandleContext(context)
	})
	router.GET("/test5", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	router.Run(":8080")
}
