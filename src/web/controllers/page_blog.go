package controllers

import (
	"github.com/astaxie/beego"
)

type Page_blog struct {
	beego.Controller
}

func (c *Page_blog) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "page_blog.html"
}
