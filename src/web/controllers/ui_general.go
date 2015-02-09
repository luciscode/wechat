package controllers

import (
	"github.com/astaxie/beego"
)

type Ui_general struct {
	beego.Controller
}

func (c *Ui_general) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "ui_general.html"
}
