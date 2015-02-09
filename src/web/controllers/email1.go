package controllers

import (
	"github.com/astaxie/beego"
)

type Email1 struct {
	beego.Controller
}

func (c *Email1) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "email1.html"
}
