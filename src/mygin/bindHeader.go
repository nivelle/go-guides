package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		h := testHeader{}

		if err := context.ShouldBindHeader(&h); err != nil {
			context.JSON(http.StatusOK, err)
		}
		fmt.Sprintf("%#v\n", h)
		context.JSON(http.StatusOK, gin.H{
			"Rate": h.Rate, "Domain": h.Domain})

	})
	r.Run(":8090")
}
