package controllers

import (
	"github.com/astaxie/beego"
)

type Table_basic struct {
	beego.Controller
}

func (c *Table_basic) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "table_basic.html"
}
