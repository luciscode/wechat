package controllers

import (
	"github.com/astaxie/beego"
)

type Layout_horizontal_menu1 struct {
	beego.Controller
}

func (c *Layout_horizontal_menu1) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "layout_horizontal_menu1.html"
}
