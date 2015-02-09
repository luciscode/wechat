package controllers

import (
	"github.com/astaxie/beego"
)

type Test struct {
	beego.Controller
}

func (c *Test) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "test.html"
}
