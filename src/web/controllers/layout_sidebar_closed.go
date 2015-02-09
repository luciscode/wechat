package controllers

import (
	"github.com/astaxie/beego"
)

type Layout_sidebar_closed struct {
	beego.Controller
}

func (c *Layout_sidebar_closed) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "layout_sidebar_closed.html"
}
