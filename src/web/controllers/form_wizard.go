package controllers

import (
	"github.com/astaxie/beego"
)

type Form_wizard struct {
	beego.Controller
}

func (c *Form_wizard) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "form_wizard.html"
}
