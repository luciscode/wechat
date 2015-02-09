package controllers

import (
	"github.com/astaxie/beego"
)

type Page_timeline struct {
	beego.Controller
}

func (c *Page_timeline) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "page_timeline.html"
}
