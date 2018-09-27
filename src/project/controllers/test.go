package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (c *TestController) Index() {

}

func (c *TestController) List() {
	c.TplName = "test/index.tpl"
}