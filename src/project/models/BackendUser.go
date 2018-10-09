package models

import "github.com/astaxie/beego/orm"

func (a *BackendUser) TableName() string {
	return BackendUserTBName()
}

type BackendUser struct {
	Id                 int
	RealName           string `orm:"size(32)"`
	UserName           string `orm:"size(24)"`
	UserPwd            string `json:"-"`
	IsSuper            bool
	Status             int
	Mobile             string                `orm:"size(16)"`
	Email              string                `orm:"size(256)"`
	Avatar             string                `orm:"size(256)"`
	RoleIds            []int                 `orm:"-" form:"RoleIds"`
	//RoleBackendUserRel []*RoleBackendUserRel `orm:"reverse(many)"` // 设置一对多的反向关系
	//ResourceUrlForList []string              `orm:"-"`
}

//验证用户名和密码
func BackendUserOneByUserName(name string,password string) (*BackendUser,error) {
	o := orm.NewOrm()
	m := BackendUser{}
	err := o.QueryTable(BackendUserTBName()).Filter("UserName",name).Filter("UserPwd",password).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

//获取用户信息
func GetUserInfoByUserId(id int)(*BackendUser,error)  {
	o := orm.NewOrm()
	m := BackendUser{Id:id}
	err := o.Read(&m)
	if err != nil{
		return nil,err
	}
	return &m,err
}
