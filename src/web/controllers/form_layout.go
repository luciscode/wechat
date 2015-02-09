package controllers

import (
	"github.com/astaxie/beego"
)

type Form_layout struct {
	beego.Controller
}

func (c *Form_layout) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "form_layout.html"
}
