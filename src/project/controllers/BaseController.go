package controllers

import (
	"github.com/astaxie/beego"
	"project/models"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string             //当前控制名称
	actionName     string             //当前action名称
	curUser        models.BackendUser //当前用户信息
}

func (c *BaseController) Prepare() {
	//附值
	//c.controllerName, c.actionName = c.GetControllerAndAction()
	//从Session里获取数据 设置用户信息
	c.adapterUserInfo()
}

//从session里取用户信息
func (c *BaseController) adapterUserInfo() {
	a := c.GetSession("backenduser")
	if a != nil {
		c.curUser = a.(models.BackendUser)
		c.Data["backenduser"] = a
	}
}

//判断是否登录，未登录则跳转至登录页面
func (c *BaseController) checkLogin()  {
	if c.curUser.Id == 0 {
		//登录地址页面
		urlstr := c.URLFor("HomeController.Login")
		//登陆成功后，返回当前地址
		//returnURL := c.Ctx.Request.URL.Path
		//如果ajax请求则返回相应的错码和跳转的地址
		if c.Ctx.Input.IsAjax() {
			//由于是ajax请求，因此地址是header里的Referer
			returnURL := c.Ctx.Input.Refer()
			c.jsonResult(302, "请登录", urlstr+returnURL)
		}
		c.Redirect(urlstr, 302)
		c.StopRun()
	}
}

//返回json数据
func (c *BaseController) jsonResult(code int64,msg string,obj interface{})  {
	r := &models.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

// 设置模板
// 第一个参数模板，第二个参数为layout
func (c *BaseController) setTpl(template ...string) {
	var tplName string
	var layout string
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		//不要Controller这个10个字母
		ctrlName := strings.ToLower(c.controllerName[0 : len(c.controllerName)-10])
		actionName := strings.ToLower(c.actionName)
		tplName = ctrlName + "/" + actionName + ".html"
	}
	c.Layout = layout
	c.TplName = tplName
}
//获取session
func (c *BaseController)setBackendUser2Session(userId int) error  {
	m,err := models.GetUserInfoByUserId(userId)
	if err!=nil{
		return err
	}
	c.SetSession("backenduser", *m)
	return nil
}

//// 重定向
//func (c *BaseController) pageUrl(template string) {
//	url := c.URLFor(template)
//	c.Redirect(url, 302)
//	c.StopRun()
//}

// 重定向
func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}
