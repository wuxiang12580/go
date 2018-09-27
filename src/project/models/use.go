package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id int
	Name string
}

func init() {
	// 需要在 init 中注册定义的 model
	orm.RegisterModel(new(User))
}
