package admin

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController) Index()  {

	c.Layout = "admin/layout.html"
	c.TplName="admin/user/index.tpl"
}
