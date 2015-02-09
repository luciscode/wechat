package controllers

import (
	"github.com/astaxie/beego"
)

type Layout_ajax struct {
	beego.Controller
}

func (c *Layout_ajax) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "layout_ajax.html"
}
