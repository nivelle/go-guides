package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.GET("/", indexHandler)
	router.POST("/", formHandler)

	router.Run(":8080")
}
