package controllers

import (
	"github.com/astaxie/beego"
)

type Email4 struct {
	beego.Controller
}

func (c *Email4) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "email4.html"
}
