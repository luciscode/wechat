package controllers

import (
	"github.com/astaxie/beego"
)

type Extra_faq struct {
	beego.Controller
}

func (c *Extra_faq) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "extra_faq.html"
}
