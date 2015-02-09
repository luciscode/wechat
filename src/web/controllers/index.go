package controllers

import (
	"github.com/astaxie/beego"
)

type Index struct {
	beego.Controller
}

func (c *Index) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.html"
}
