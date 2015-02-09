package controllers

import (
	"github.com/astaxie/beego"
)

type Ui_tree struct {
	beego.Controller
}

func (c *Ui_tree) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "ui_tree.html"
}
