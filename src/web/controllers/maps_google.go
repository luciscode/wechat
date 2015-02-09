package controllers

import (
	"github.com/astaxie/beego"
)

type Maps_google struct {
	beego.Controller
}

func (c *Maps_google) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "maps_google.html"
}
