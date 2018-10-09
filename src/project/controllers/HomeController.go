package controllers

type HomeController struct {
	BaseController
}

func (c *HomeController) Index()  {
	//判断是否登录
	c.checkLogin()
}