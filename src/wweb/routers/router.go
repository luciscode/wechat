package routers

import (
	"github.com/astaxie/beego"
	"wweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/editor", &controllers.Editor{})

}
