package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	router := gin.Default()

	//在中间件或处理程序中启动新的Goroutines时，你不应该使用其中的原始上下文，你必须使用只读副本
	//（c.Copy()）
	router.GET("long_async", func(context *gin.Context) {
		cCp := context.Copy()
		go func() {
			time.Sleep(5 * time.Second)

			log.Println("done in path", cCp.Request.URL.Path)
		}()
	})

	router.GET("long_sync", func(context *gin.Context) {
		time.Sleep(5*time.Second)

		log.Println("done in path", context.Request.URL.Path)
	})
}
