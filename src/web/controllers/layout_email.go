package controllers

import (
	"github.com/astaxie/beego"
)

type Layout_email struct {
	beego.Controller
}

func (c *Layout_email) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "layout_email.html"
}
