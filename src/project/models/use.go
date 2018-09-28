package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int `orm:"pk"`
	Name string
	Password string
	Add_time int
}

func init() {
	// 需要在 init 中注册定义的 model
	orm.RegisterModel(new(User))
}

func ValidateUser(name string,password string) (User,error) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("name",name).Filter("password",password).One(&user)
	return user,err
}
