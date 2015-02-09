package controllers

import (
	"github.com/astaxie/beego"
)

type Extra_404_option1 struct {
	beego.Controller
}

func (c *Extra_404_option1) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "extra_404_option1.html"
}
