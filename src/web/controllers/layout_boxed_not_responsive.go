package controllers

import (
	"github.com/astaxie/beego"
)

type Layout_boxed_not_responsive struct {
	beego.Controller
}

func (c *Layout_boxed_not_responsive) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "layout_boxed_not_responsive.html"
}
