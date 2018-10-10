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

//获取用户总数
func CountArticle(data map[string]string) int64  {
	o := orm.NewOrm()
	qs := o.QueryTable(BackendUserTBName())
	con := orm.NewCondition()
	if data["title"] != "" {
		con = con.And("UserName",data["title"])
	}

	if data["keywords"] != "" {
		con = con.And("Mobile",data["keywords"])
	}

	num, _ := qs.SetCond(con).Count()
	return num
}

//获取列表
func ListArticle(data map[string]string,page,offset int) (num int64, err error, art []BackendUser) {
	o := orm.NewOrm()
	qs := o.QueryTable(BackendUserTBName())
	con := orm.NewCondition()
	if data["title"] != "" {
		con = con.And("UserName", data["title"])
	}
	if data["keywords"] != "" {
		con = con.Or("Mobile", data["keywords"])
	}
	qs = qs.SetCond(con)

	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset = 10
	}
	start := (page - 1) * offset
	var backendUser []BackendUser
	num, err1 := qs.Limit(offset, start).All(&backendUser)
	return num, err1, backendUser
}
