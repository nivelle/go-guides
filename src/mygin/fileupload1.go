package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"fmt"
)

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 16 << 20
	router.POST("/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		log.Println(file.Filename)
		context.SaveUploadedFile(file, "../")
		context.String(http.StatusOK, fmt.Sprintf(" %s upload", file.Filename))
	})

	router.MaxMultipartMemory = 8<<20
	router.Run(":8000")
}
