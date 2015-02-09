package controllers

import (
	"github.com/astaxie/beego"
)

type Table_managed struct {
	beego.Controller
}

func (c *Table_managed) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "table_managed.html"
}
