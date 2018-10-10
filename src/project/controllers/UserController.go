package controllers

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"project/models"
	"fmt"
)

type UserController struct {
	BaseController
}

func (c *UserController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	c.checkLogin()
}

func (c *UserController) Index()  {
	page,err := c.GetInt("p")
	if err!=nil {
		page = 1
	}
	offset,err := beego.AppConfig.Int("pageoffset")
	if err!=nil{
		offset = 10
	}
	//搜索参数
	title := strings.TrimSpace(c.GetString("username"))
	keywords := strings.TrimSpace(c.GetString("mobile"))

	condArr := make(map[string]string)
	condArr["title"] = title
	condArr["keywords"] = keywords

	countArticle := models.CountArticle(condArr)

	paginator := pagination.SetPaginator(c.Ctx, offset, countArticle)
	_, _, art := models.ListArticle(condArr, page, offset)

	c.Data["paginator"] = paginator
	c.Data["art"] = art

	c.setTpl("admin/user/index.html","admin/layout.html")
}

func (c *UserController) Add()  {
	c.setTpl("admin/user/add.html","admin/layout.html")
}
