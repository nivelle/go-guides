package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("fuck !!!!"+beego.AppConfig.String("myLoveName"))
}

func (c *UserController) ShowInfo() {
	info := "my info is great !!"
	fmt.Print("my info is great !!")
	c.Ctx.WriteString(info)
}
