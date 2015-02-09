package controllers

import (
	"github.com/astaxie/beego"
)

type Layout_horizontal_menu2 struct {
	beego.Controller
}

func (c *Layout_horizontal_menu2) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "layout_horizontal_menu2.html"
}
