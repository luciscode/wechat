package controllers

import (
	"github.com/astaxie/beego"
)

type Layout_promo struct {
	beego.Controller
}

func (c *Layout_promo) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "layout_promo.html"
}
