package controllers

import (
	"github.com/astaxie/beego"
)

type Form_dropzone struct {
	beego.Controller
}

func (c *Form_dropzone) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "form_dropzone.html"
}
