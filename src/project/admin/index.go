package admin

import (
	"github.com/astaxie/beego"
	"project/models"
	"fmt"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Index()  {
	sess := c.StartSession()
	sessionUid := sess.Get("uid")
	if sessionUid==nil{
		c.Redirect("/login/index", 302)
	}
	var user models.User
	user.Id = sessionUid.(int)

	userInfo,err := models.GetUserInfoById(user.Id)
	if err!=nil{
		fmt.Println("用户不存在")
	}
	c.Data["user"]=userInfo
	//布局页面
	c.Layout = "admin/layout.html"
	//内容页面
	c.TplName = "admin/index/index.tpl"
}