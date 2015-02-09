package controllers

import (
	"github.com/astaxie/beego"
)

type Extra_lock struct {
	beego.Controller
}

func (c *Extra_lock) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "extra_lock.html"
}
