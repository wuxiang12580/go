package admin

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Index()  {
	session := c.StartSession()
	uid := session.Get("uid")
	if uid==nil{
		c.Redirect("/login/index", 302)
	}else{
		c.Redirect("/index/index", 302)
	}
}
