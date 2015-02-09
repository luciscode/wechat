package controllers

import (
	"github.com/astaxie/beego"
)

type Form_fileupload struct {
	beego.Controller
}

func (c *Form_fileupload) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "form_fileupload.html"
}
