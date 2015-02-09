package controllers

import (
	"github.com/astaxie/beego"
)

type Autoreply struct {
	beego.Controller
}

func (c *Autoreply) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "autoreply.html"
}
