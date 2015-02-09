package controllers

import (
	"github.com/astaxie/beego"
)

type Extra_404_option2 struct {
	beego.Controller
}

func (c *Extra_404_option2) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "extra_404_option2.html"
}
