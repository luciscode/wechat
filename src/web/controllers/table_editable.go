package controllers

import (
	"github.com/astaxie/beego"
)

type Table_editable struct {
	beego.Controller
}

func (c *Table_editable) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "table_editable.html"
}
