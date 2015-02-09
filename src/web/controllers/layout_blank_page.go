package controllers

import (
	"github.com/astaxie/beego"
)

type Layout_blank_page struct {
	beego.Controller
}

func (c *Layout_blank_page) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "layout_blank_page.html"
}
