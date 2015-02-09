package controllers

import (
	"github.com/astaxie/beego"
)

type Head struct {
	beego.Controller
}

func (c *Head) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "head.html"
}
