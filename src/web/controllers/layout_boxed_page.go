package controllers

import (
	"github.com/astaxie/beego"
)

type Layout_boxed_page struct {
	beego.Controller
}

func (c *Layout_boxed_page) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "layout_boxed_page.html"
}
