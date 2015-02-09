package controllers

import (
	"github.com/astaxie/beego"
)

type Layout_horizontal_sidebar_menu struct {
	beego.Controller
}

func (c *Layout_horizontal_sidebar_menu) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "layout_horizontal_sidebar_menu.html"
}
