package controllers

import (
	"github.com/astaxie/beego"
)

type Addautoreplynews struct {
	beego.Controller
}

func (c *Addautoreplynews) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "addautoreplynews.html"
}
