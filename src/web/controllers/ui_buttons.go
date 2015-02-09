package controllers

import (
	"github.com/astaxie/beego"
)

type Ui_buttons struct {
	beego.Controller
}

func (c *Ui_buttons) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "ui_buttons.html"
}
