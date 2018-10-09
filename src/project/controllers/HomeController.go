package controllers

type HomeController struct {
	BaseController
}

func (c *HomeController) Index()  {
	//判断是否登录
	c.checkLogin()
	c.setTpl()
}

func (c *HomeController) Login()  {
	c.TplName = "admin/login/index.tpl"
}