package controllers

import (
	"github.com/astaxie/beego"
	_"github.com/astaxie/beego/orm"
	_"beego1/models"

	"github.com/astaxie/beego/orm"
	"beego1/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["data"]="babyhehehe"
	c.TplName = "test.html"
}
func (c *MainController) Post() {
	c.Data["data"]="LOVE"
	c.TplName = "test.html"
}
func (c *MainController) ShowGet() {

    /* o:=orm.NewOrm()

	var USE models.Test

	USE.Name="LILYhhhhhhhhhhhh"
	USE.PassWord="123456"

	cont,err:=o.Insert(&USE)
	if err!=nil{
		beego.Error("插入失败")
	}
	beego.Info(cont)*/
	/*o:=orm.NewOrm()
	var use models.Test

	use.Id=1
	err:=o.Read(&use,"id")

	if err!=nil{
		beego.Error("查询失败")
	}
	beego.Info(use)*/
	/*
	o:=orm.NewOrm()
	var use models.Test

	use.Id=1

	err:=o.Read(&use)
	 if err!=nil{
	 	beego.Error("更新的数据不存在")
	 }
	 use.Name="girl"
     use.PassWord="8866688"
     cont,err:=o.Update(&use)
     if err!=nil{
     	beego.Error("更新失败")
	 }
	 beego.Info(cont)
	*/
    o:=orm.NewOrm()
    var use models.Test

    use.Id=1

    cont,err:=o.Delete(&use)
     if err!=nil{
     	beego.Error("删除失败")
	 }
	 beego.Info(cont)



	c.Data["data"]="LOVE"
	c.TplName = "test.html"
}
