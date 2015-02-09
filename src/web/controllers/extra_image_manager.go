package controllers

import (
	"github.com/astaxie/beego"
)

type Extra_image_manager struct {
	beego.Controller
}

func (c *Extra_image_manager) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "extra_image_manager.html"
}
