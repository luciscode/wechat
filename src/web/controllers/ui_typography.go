package controllers

import (
	"github.com/astaxie/beego"
)

type Ui_typography struct {
	beego.Controller
}

func (c *Ui_typography) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "ui_typography.html"
}
