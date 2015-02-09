package controllers

import (
	"github.com/astaxie/beego"
)

type Ui_sliders struct {
	beego.Controller
}

func (c *Ui_sliders) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "ui_sliders.html"
}
