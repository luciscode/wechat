package controllers

import (
	"github.com/astaxie/beego"
)

type Inbox struct {
	beego.Controller
}

func (c *Inbox) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "inbox.html"
}
