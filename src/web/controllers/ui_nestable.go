package controllers

import (
	"github.com/astaxie/beego"
)

type Ui_nestable struct {
	beego.Controller
}

func (c *Ui_nestable) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "ui_nestable.html"
}
