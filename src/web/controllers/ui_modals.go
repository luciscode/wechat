package controllers

import (
	"github.com/astaxie/beego"
)

type Ui_modals struct {
	beego.Controller
}

func (c *Ui_modals) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "ui_modals.html"
}
