package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"-"` //- 跳过验证
}

func main() {

	router := gin.Default()
	/**
	类型 - Should bind
    方法 - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery,ShouldBindYAML, ShouldBindHeader。
    行为 - 这些方法底层使用 ShouldBindWith ，如果存在绑定错误，则返回错误，开发人员可以正确处理请求和错误。当我们使用绑定方法时，Gin会根据 Content-Type 推断出使用哪种绑定器，如果你确定你绑定的是什么，你可以使用 MustBindWith 或者 BindingWith 。
	*/
	router.POST("/loginJSON", func(context *gin.Context) {
		var json Login
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if json.User != "nivelle" || json.Password != "jessy" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized"})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"status": "logged in"})
	})
	router.Run(":9001")
}
