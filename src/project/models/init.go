package models

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

func init()  {
	orm.RegisterModel(new(BackendUser))
}

//下面是统一的表名管理
func TableName(name string) string  {
	prefix := beego.AppConfig.String("db_prefix")
	return prefix + name
}

//分别获取对应表的名称
func BackendUserTBName()string  {
	return TableName("backend_user")
}