package controllers

import (
	"github.com/astaxie/beego"
)

type Email6 struct {
	beego.Controller
}

func (c *Email6) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "email6.html"
}
