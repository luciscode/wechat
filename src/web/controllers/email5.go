package controllers

import (
	"github.com/astaxie/beego"
)

type Email5 struct {
	beego.Controller
}

func (c *Email5) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "email5.html"
}
