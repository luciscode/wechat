package controllers

import (
	"github.com/astaxie/beego"
)

type Page_calendar struct {
	beego.Controller
}

func (c *Page_calendar) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "Page_calendar.html"
}
