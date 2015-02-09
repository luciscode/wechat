package controllers

import (
	"github.com/astaxie/beego"
)

type Email2 struct {
	beego.Controller
}

func (c *Email2) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "email2.html"
}
