package controllers

import (
	"github.com/astaxie/beego"
)

type Form_validation struct {
	beego.Controller
}

func (c *Form_validation) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "form_validation.html"
}
