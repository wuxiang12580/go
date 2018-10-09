package controllers

import (
	"strings"
	"project/util"
	"project/models"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index()  {
	//判断是否登录
	c.checkLogin()
	c.setTpl("admin/home/index.html","admin/layout.html")
}

func (c *HomeController) Login()  {
	c.setTpl("admin/login/index.tpl")
}

func (c *HomeController) DoLogin()  {
	username := strings.TrimSpace(c.GetString("username"))
	pwd := strings.TrimSpace(c.GetString("password"))

	if username =="" || pwd == ""{
		c.jsonResult(1,"用户名和密码不正确","")
	}
	//密码MD5加密
	pwd = util.Str2md5(pwd)
	//是否存在用户名和密码
	user,err := models.BackendUserOneByUserName(username,pwd)
	if user != nil && err == nil {
		//保存用户信息到session
		c.setBackendUser2Session(user.Id)
		//获取用户信息
		c.jsonResult(4, "登录成功", "")
	} else {
		c.jsonResult(3, "用户名或者密码错误", "")
	}

}

func (c *HomeController) Logout() {
	user := models.BackendUser{}
	c.SetSession("backenduser", user)
	c.pageLogin()
}