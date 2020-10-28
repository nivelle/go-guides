package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20

	router.POST("/upload", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			println(file.Filename)
			context.SaveUploadedFile(file, "../")
		}
		context.String(http.StatusOK, fmt.Sprintf("%d files uploaded", len(files)))
	})
	router.Run(":8000")

}
