package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	router := gin.Default()

	router.GET("/someJSON", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "key", "status": http.StatusOK})
	})

	router.GET("moreJSON", func(context *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		context.JSON(http.StatusOK, msg)
	})

	router.GET("/someXML", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	//yaml文件解析
	router.GET("/someYAML", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.GET("someProtoBuf", func(context *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		context.ProtoBuf(http.StatusOK, data)
	})
	router.Run(":8080")
}
