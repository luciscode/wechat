package controllers

import (
	"github.com/astaxie/beego"
)

type Layout_ajax_content_2 struct {
	beego.Controller
}

func (c *Layout_ajax_content_2) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "layout_ajax_content_2.html"
}
