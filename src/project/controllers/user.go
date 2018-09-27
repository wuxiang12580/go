package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"project/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Index()  {
	o := orm.NewOrm()
	user := new(models.User)
	user.Name = "slene"
	fmt.Println(o.Insert(user))
}

func (c *UserController) Get()  {
	id,_ := c.GetInt(":id")
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	res := qs.Filter("id", id)
	//user := models.User{Id: id}
	//err := o.Read(&user)
	fmt.Println(res)
}

func (c *UserController) Delete()  {
	id,_ := c.GetInt(":id")
	o := orm.NewOrm()
	user := models.User{Id: id}
	num , err := o.Delete(&user)
	if err==nil{
		fmt.Println(num)
	}
}
