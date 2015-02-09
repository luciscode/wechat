package controllers

import (
	"github.com/astaxie/beego"
)

type Extra_search struct {
	beego.Controller
}

func (c *Extra_search) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "extra_search.html"
}
