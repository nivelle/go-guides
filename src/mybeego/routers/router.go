package routers

import (
	"mybeego/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//一：固定路由
    beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
    //指定具体的方法
    /**
    1. * 表示任意的 method 都执行该函数
    2. 使用 httpmethod:funcname 格式来展示
    3. 多个不同的格式使用 ; 分割
    4. 多个 method 对应同一个 funcname，method 之间通过 , 来分割
     */
	beego.Router("/user/info",&controllers.UserController{},"get:ShowInfo")

    //二：基础路由
    beego.Any("/user/detail", func(context *context.Context) {
		context.Output.Body([]byte("bar"))
	})

	//三：自动匹配
    //http://localhost:8080/auto/login
	beego.AutoRouter(&controllers.AutoController{})


}


