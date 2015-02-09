package controllers

import (
	"github.com/astaxie/beego"
)

type Form_component struct {
	beego.Controller
}

func (c *Form_component) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "form_component.html"
}
