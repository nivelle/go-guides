package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 默认已经连接了 Logger and Recovery 中间件
	//router := gin.Default()

	//没有中间件的默认路由
	router := gin.New()
	//全局中间件
	// Logger 中间件将写日志到 gin.DefaultWriter 即使你设置 GIN_MODE=release
	// 默认设置 gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	//Recovery 中间件从任何 panic 恢复，如果出现 panic，它会写一个 500 错误
	router.Use(gin.Recovery())
	// 对于每个路由中间件，您可以根据需要添加任意数量
	router.GET("/benchmark", func(context *gin.Context) {

	}, func(context *gin.Context) {

	})

	authorized := router.Group("/")
	{
		authorized.POST("/test", func(context *gin.Context) {
			context.String(http.StatusOK, "test is ok")
		})
		// 嵌套组
		testing := authorized.Group("testing")
		testing.GET("/inner", func(context *gin.Context) {
			context.String(http.StatusOK, "inner group is ok")
		})
	}
	router.Run(":8000")

}
