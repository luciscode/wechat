package controllers

import (
	"github.com/astaxie/beego"
)

type Page_coming_soon struct {
	beego.Controller
}

func (c *Page_coming_soon) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "page_coming_soon.html"
}
