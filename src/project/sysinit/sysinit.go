package sysinit

import "github.com/astaxie/beego"

func init()  {
	//启用session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//初始化日志

	//初始化缓存

	//初始化数据库
	InitDatabase()
}