package controllers

import (
	"github.com/astaxie/beego"
)

type All struct {
	beego.Controller
}

func (c *All) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "all.html"
}
