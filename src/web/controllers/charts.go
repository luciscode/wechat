package controllers

import (
	"github.com/astaxie/beego"
)

type Charts struct {
	beego.Controller
}

func (c *Charts) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "charts.html"
}
