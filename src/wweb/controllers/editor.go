package controllers

import (
	"github.com/astaxie/beego"
)

type Editor struct {
	beego.Controller
}

func (c *Editor) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "editor.html"
}
