package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person3 struct {
	Name string `uri:"name" binding:"required"`
	Id   string `uri:"id" binding:"required,uuid"`
}

func main() {
	router := gin.Default()
	router.GET("/:name/:id", func(context *gin.Context) {
		var person Person3
		if err := context.ShouldBindUri(&person); err != nil {
			context.JSON(400, gin.H{"msg": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.Id})
	})
	router.Run(":8088")
}
