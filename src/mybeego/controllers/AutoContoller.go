package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type AutoController struct {
	beego.Controller
	name string
}

func (c *AutoController) Login() {
	c.Ctx.WriteString("AutoController get")
}

func (c *AutoController) Logout() {
	info := "AutoController post"
	fmt.Print("my post is great !!")
	c.Ctx.WriteString(info)
}