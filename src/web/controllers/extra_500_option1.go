package controllers

import (
	"github.com/astaxie/beego"
)

type Extra_500_option1 struct {
	beego.Controller
}

func (c *Extra_500_option1) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "extra_500_option1.html"
}
