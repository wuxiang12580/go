package controllers

type UserCenterController struct {
	BaseController
}

func (c *UserCenterController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	c.checkLogin()
}

func (c *UserCenterController) Index()  {
	c.setTpl("admin/usercenter/index.html","admin/layout.html")
}

func (c *UserCenterController) Add()  {
	c.setTpl("admin/usercenter/edit.html","admin/layout.html")
}
