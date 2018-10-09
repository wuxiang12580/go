package models

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
