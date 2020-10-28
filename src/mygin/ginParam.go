package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	router := gin.Default()

	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "hello %s", name)
	})

	router.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + " is " + action
		context.String(http.StatusOK, message)
	})
	//对于每个匹配的请求，上下文将保留路由定义
	router.POST("/user/:name/*action", func(context *gin.Context) {
		result := context.FullPath() == "/user/:name/*action"
		fmt.Println(result)
	})
	// query?param1=value1&param2=value2
	router.GET("/welcome", func(context *gin.Context) {
		// 默认值 为nivelle
		firstName := context.DefaultQuery("firstName", "nivelle")
		lastName := context.Query("lastName")
		context.String(http.StatusOK, "hello %s,%s", firstName, lastName)
	})
	//post表单提交
	router.POST("/myForm", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "fuck")
		context.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.POST("/post", func(context *gin.Context) {
		ids:= context.QueryMap("ids")
		names:=context.PostFormMap("names")
		fmt.Printf("ids: %v,names:%v",ids,names)
	})

	router.Run(":8000")
}
