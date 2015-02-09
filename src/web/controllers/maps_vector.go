package controllers

import (
	"github.com/astaxie/beego"
)

type Maps_vector struct {
	beego.Controller
}

func (c *Maps_vector) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "maps_vector.html"
}
