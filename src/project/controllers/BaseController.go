package controllers

import (
	"github.com/astaxie/beego"
	"project/models"
)

type BaseController struct {
	beego.Controller
	curUser        models.BackendUser //当前用户信息
}

//判断是否登录，未登录则跳转至登录页面
func (c *BaseController) checkLogin()  {
	if c.curUser.Id == 0 {
		//登录地址页面
		urlstr := c.URLFor("HomeController.Login") + "?url="
		//登陆成功后，返回当前地址
		returnURL := c.Ctx.Request.URL.Path
		//如果ajax请求则返回相应的错码和跳转的地址
		if c.Ctx.Input.IsAjax() {
			//由于是ajax请求，因此地址是header里的Referer
			returnURL := c.Ctx.Input.Refer()
			c.jsonResult(enums.JRCode302, "请登录", urlstr+returnURL)
		}
		c.Redirect(urlstr+returnURL, 302)
		c.StopRun()
	}
}

func (c *BaseController) jsonResult(code int64,msg string,)  {

}
