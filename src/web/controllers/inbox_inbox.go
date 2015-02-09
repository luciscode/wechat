package controllers

import (
	"github.com/astaxie/beego"
)

type Inbox_inbox struct {
	beego.Controller
}

func (c *Inbox_inbox) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "inbox_inbox.html"
}
