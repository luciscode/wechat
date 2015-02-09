package controllers

import (
	"github.com/astaxie/beego"
)

type Page_about struct {
	beego.Controller
}

func (c *Page_about) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "page_about.html"
}
