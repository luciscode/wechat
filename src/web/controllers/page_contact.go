package controllers

import (
	"github.com/astaxie/beego"
)

type Page_contact struct {
	beego.Controller
}

func (c *Page_contact) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "page_contact.html"
}
