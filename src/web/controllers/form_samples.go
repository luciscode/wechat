package controllers

import (
	"github.com/astaxie/beego"
)

type Form_samples struct {
	beego.Controller
}

func (c *Form_samples) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "form_samples.html"
}
