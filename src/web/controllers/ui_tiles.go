package controllers

import (
	"github.com/astaxie/beego"
)

type Ui_tiles struct {
	beego.Controller
}

func (c *Ui_tiles) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "ui_tiles.html"
}
