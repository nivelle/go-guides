package main

import (
	"mime/multipart"
	"github.com/gin-gonic/gin"
	"net/http"
)

type profileForm struct {
	Name   string                `form:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

func main() {
	router := gin.Default()

	router.POST("/profile", func(context *gin.Context) {
		var form profileForm

		if err := context.ShouldBind(&form); err != nil {
			context.String(http.StatusBadRequest, "bad request")
			return
		}
		err := context.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
		if err != nil {
			context.String(http.StatusInternalServerError, "unknown error")
			return
		}
		context.String(http.StatusOK, "ok")
	})
	router.Run(":8080")
}
