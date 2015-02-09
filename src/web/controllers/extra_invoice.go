package controllers

import (
	"github.com/astaxie/beego"
)

type Extra_invoice struct {
	beego.Controller
}

func (c *Extra_invoice) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "extra_invoice.html"
}
