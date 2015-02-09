package controllers

import (
	"github.com/astaxie/beego"
)

type Page_blog_item struct {
	beego.Controller
}

func (c *Page_blog_item) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "page_blog_item.html"
}
