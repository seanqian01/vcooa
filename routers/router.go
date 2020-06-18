package routers

import (
	"github.com/astaxie/beego"
	"vcooa/controllers"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/admin", &controllers.MainController{})
	beego.Router("/com", &controllers.ComController{})
}
