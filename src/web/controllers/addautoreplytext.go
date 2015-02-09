package controllers

import (
	"github.com/astaxie/beego"
)

type Addautoreplytext struct {
	beego.Controller
}

func (c *Addautoreplytext) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "addautoreplytext.html"
}
