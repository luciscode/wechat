package controllers

import (
	"github.com/astaxie/beego"
)

type Table_advanced struct {
	beego.Controller
}

func (c *Table_advanced) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "table_advanced.html"
}
