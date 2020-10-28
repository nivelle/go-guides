package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	{
		v1.POST("/test1", func(context *gin.Context) {
			context.String(http.StatusOK, "v1/test1 is ok")
		})
	}

	v2 := router.Group("/v2")
	{
		v2.POST("test1", func(context *gin.Context) {
			context.String(http.StatusOK, "v2/test1 is ok")

		})
	}

	router.Run(":8000")
}
