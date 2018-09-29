package routers

import (
	"github.com/astaxie/beego"
	"project/admin"
)

func init() {
    /*beego.Router("/", &controllers.MainController{})
	beego.Router("/user/index", &controllers.UserController{},"*:Index")
	beego.Router("/user/get/:id([0-9]+)", &controllers.UserController{},"*:Get")
	beego.Router("/user/del/:id([0-9]+)", &controllers.UserController{},"*:Delete")
	beego.Router("/test/index", &controllers.TestController{},"*:Index")
	beego.Router("/api/list", &controllers.TestController{}, "*:List")
	*/
	beego.Router("/admin", &admin.AdminController{},"*:Index")
	beego.Router("/login/index", &admin.LoginController{},"*:Index")
	beego.Router("/login/login", &admin.LoginController{},"*:Login")
	beego.Router("/login/logout", &admin.LoginController{},"*:Logout")
	beego.Router("/index/index", &admin.IndexController{},"*:Index")

	beego.Router("/user/index", &admin.UserController{},"*:Index")
	beego.Router("/user/add", &admin.UserController{},"*:Add")
}
