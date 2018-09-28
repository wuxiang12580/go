package admin

import (
	"github.com/astaxie/beego"
	"fmt"
	"project/models"
)

type LoginController struct {
	beego.Controller
}
func (c *LoginController) Index() {
	c.TplName = "admin/login/index.tpl"
}

func (c *LoginController) Login()  {
	sess := c.StartSession()
	name := c.Input().Get("username")
	password := c.Input().Get("password")
	data,err := models.ValidateUser(name,password);
	if err!=nil{
		fmt.Println("用户名或密码错误！")
		c.Redirect("/login/index", 302)
		return
	}
	sess.Set("uid", data.Id)
	c.Redirect("/index/index", 302)
}

func (c *LoginController) Logout()  {
	if c.Ctx.Request.Method == "GET" {
		sess := c.StartSession()
		sess.Delete("uid")
		c.Redirect("/login/index", 302)
	}
}
