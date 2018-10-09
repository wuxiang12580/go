package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "project/models"
)

//初始化数据库连接
func InitDatabase()  {
	//读取配置文件，设置数据库参数
	//数据库类别
	dbType := beego.AppConfig.String("db_type")
	//数据库连接名称
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	//数据库名称
	dbName := beego.AppConfig.String(dbType + "::db_name")
	//数据库连接用户名
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	//数据库连接密码
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	//数据库ip
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	//数据库端口
	dbPort := beego.AppConfig.String(dbType + "::db_port")

	dbCharset := beego.AppConfig.String(dbType + "::db_charset")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(dbAlias, dbType, dbUser+":"+dbPwd+"@tcp("+dbHost+":"+
		dbPort+")/"+dbName+"?charset="+dbCharset, 30)

	//如果是开发模式，则显示命令信息
	isDev := beego.AppConfig.String("runmode")
	if isDev=="dev" {
		orm.Debug = true
	}
	orm.RunSyncdb("default", false, true)
}