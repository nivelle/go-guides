package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "nivelle.me"
	c.Data["Email"] = "fuxinzhong123@gmail.com"
	c.TplName = "index.html"
}
