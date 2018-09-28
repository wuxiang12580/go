package admin

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Index()  {
	//布局页面
	c.Layout = "admin/layout.html"
	//内容页面
	c.TplName = "admin/index/index.tpl"
}