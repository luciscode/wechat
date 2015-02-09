package controllers

import (
	"github.com/astaxie/beego"
)

type Portlet_general struct {
	beego.Controller
}

func (c *Portlet_general) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "portlet_general.html"
}
