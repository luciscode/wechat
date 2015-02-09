package controllers

import (
	"github.com/astaxie/beego"
)

type Ui_jqueryui struct {
	beego.Controller
}

func (c *Ui_jqueryui) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "ui_jqueryui.html"
}
