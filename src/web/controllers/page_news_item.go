package controllers

import (
	"github.com/astaxie/beego"
)

type Page_news_item struct {
	beego.Controller
}

func (c *Page_news_item) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "page_news_item.html"
}
