package admin

import (
	"github.com/astaxie/beego"
	"fmt"
)

type LoginController struct {
	beego.Controller
}
func (c *LoginController) Index() {
	fmt.Println(1234)
	c.TplName = "admin/login/index.tpl"
}