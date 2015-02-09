package controllers

import (
	"github.com/astaxie/beego"
)

type Table_responsive struct {
	beego.Controller
}

func (c *Table_responsive) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "table_responsive.html"
}
