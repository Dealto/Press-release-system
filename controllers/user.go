package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"beego1/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController)ShowRegister()  {
	this.TplName="register.html"
}
func (this *UserController)HandlePost()  {
	//获得数据
	usename:=this.GetString("userName")
	pwd:=this.GetString("password")

	//检验数据
	if usename==""||pwd==""{
		this.Data["errmsg"]="注册数据不完整，请重新注册"
		this.TplName="register.html"
		return
	}
	//操作数据
	o:=orm.NewOrm()
	var use models.Test
	use.Name=usename
	use.PassWord=pwd
	o.Insert(&use)

	//返回视图
	this.Redirect("/login",302)
}
func (this *UserController)ShowLogin()  {
	userName:=this.Ctx.GetCookie("userName")
	if userName==""{
		this.Data["userName"]=""
		this.Data["checked"]=""
	}else {
		this.Data["userName"]=userName
		this.Data["checked"]="checked"
	}
	this.TplName="login.html"
}
func (this *UserController)HandleLogin() {
	//获得数据
      name:=this.GetString("userName")
      pwd:=this.GetString("password")
	//检验数据
      if name =="" || pwd ==""{
      	this.Data["errmsg"]="登陆数据不完整"
		  this.TplName="login.html"
		  return
	  }
	//操作数据并核对
	  o:=orm.NewOrm()
	  var use models.Test
	  use.Name=name
	  err:=o.Read(&use,"Name")
	  if err!=nil{
		  this.Data["errmsg"]="用户名不存在"
		  this.TplName="login.html"
		  return
	  }
	  if use.PassWord!=pwd {
		  this.Data["errmsg"]="密码错误"
		  this.TplName="login.html"
		  return
	  }
		data:=this.GetString("remember")
		beego.Info(data)
		if data=="on"{
			this.Ctx.SetCookie("userName",name,100)
		}else {
			this.Ctx.SetCookie("userName",name,-1)
		}

	  this.SetSession("username",name)
	  this.Redirect("/article/showArticlelist",302)

}

func (this *UserController)  LogOut(){
	this.DelSession("username")
	this.Redirect("/login",302)
}
