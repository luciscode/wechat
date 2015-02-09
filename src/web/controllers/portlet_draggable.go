package controllers

import (
	"github.com/astaxie/beego"
)

type Portlet_draggable struct {
	beego.Controller
}

func (c *Portlet_draggable) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "portlet_draggable.html"
}
