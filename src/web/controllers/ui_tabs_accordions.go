package controllers

import (
	"github.com/astaxie/beego"
)

type Ui_tabs_accordions struct {
	beego.Controller
}

func (c *Ui_tabs_accordions) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "ui_tabs_accordions.html"
}
