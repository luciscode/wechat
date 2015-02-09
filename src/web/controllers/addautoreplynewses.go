package controllers

import (
	"github.com/astaxie/beego"
)

type Addautoreplynewses struct {
	beego.Controller
}

func (c *Addautoreplynewses) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "addautoreplynewses.html"
}
