package controllers

import (
	"github.com/astaxie/beego"
)

type Login_soft struct {
	beego.Controller
}

func (c *Login_soft) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "login_soft.html"
}
