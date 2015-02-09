package controllers

import (
	"github.com/astaxie/beego"
)

type Extra_pricing_table struct {
	beego.Controller
}

func (c *Extra_pricing_table) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "extra_pricing_table.html"
}
