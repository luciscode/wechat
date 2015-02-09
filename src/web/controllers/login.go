package controllers

import (
	"github.com/astaxie/beego"
)

type Login struct {
	beego.Controller
}

func (c *Login) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "login.html"
}
