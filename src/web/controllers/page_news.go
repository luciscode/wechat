package controllers

import (
	"github.com/astaxie/beego"
)

type Page_news struct {
	beego.Controller
}

func (c *Page_news) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "page_news.html"
}
