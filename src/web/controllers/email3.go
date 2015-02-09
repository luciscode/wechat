package controllers

import (
	"github.com/astaxie/beego"
)

type Email3 struct {
	beego.Controller
}

func (c *Email3) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "email3.html"
}
