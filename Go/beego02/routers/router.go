package routers

import (
	"beego02/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ApiController{})
	// beego.Router("/:id", &controllers.ApiController{})
	// beego.Router("/cms_:id([0-9]+).html", &controllers.CmsController{}) //路由伪静态
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/doLogin", &controllers.LoginController{}, "post:DoLogin")
}
