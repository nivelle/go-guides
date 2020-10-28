package main

import (
	"time"
	"github.com/gin-gonic/gin"
	"log"
)

type Person2 struct {
	Name    string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"
time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func main() {
	router := gin.Default()
	router.GET("/testing", startPage1)
	router.Run(":8085")
}

func startPage1(c *gin.Context) {
	var person Person2
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses
	//`Form` (`form-data`).
	// See more at https://github.com/gingonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		log.Println(person.CreateTime)
		log.Println(person.UnixTime)
	}
	c.String(200, "Success")
}