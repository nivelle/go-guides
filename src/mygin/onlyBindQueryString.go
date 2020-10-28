package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {

	router := gin.Default()
	router.Any("testing", startPage)
	router.Run(":9002")
}

func startPage(c *gin.Context) {
	var person Person

	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(http.StatusOK, "success")
}
