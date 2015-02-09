package controllers

import (
	"github.com/astaxie/beego"
)

type Extra_profile struct {
	beego.Controller
}

func (c *Extra_profile) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "extra_profile.html"
}
