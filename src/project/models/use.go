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

//验证用户名和密码
func ValidateUser(name string,password string) (User,error) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("name",name).Filter("password",password).One(&user)
	return user,err
}
//查出用户相关信息
func GetUserInfoById(id int)(v *User,err error)  {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.QueryTable(new(User)).Filter("id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}


