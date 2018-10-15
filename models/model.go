package models

import (
	_"github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
	"time"

)
type Test struct {
	Id int
	Name string
	PassWord string
    Articles []*Article `orm:"reverse(many)"`
}
type Article struct {
	Id int `orm:"pk;auto"`
	ArtiName string `orm:"size(20)"`
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null"`
	Acontent string `orm:"size(500)"`
	Aimg string  `orm:"size(100)"`

	ArticleType *ArticleType `orm:"rel(fk)"`
	User []*Test `orm:"rel(m2m)"`
}
type ArticleType struct {
	Id int
	TypeName string `orm:"size(20)"`
	Articles []*Article `orm:"reverse(many)"`
}


func init()  {
/*
	conn,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/beegotest?charset=utf8")
     if err!=nil{
     	beego.Info("链接错误",err)
     	beego.Error("链接错误",err)
     	return
	 }
	defer conn.Close()

*//*	_,err=conn.Exec("create table user(name VARCHAR(20) ,password VARCHAR(20) );")
	if err!=nil{
		beego.Error("创建表失败",err)
		beego.Info("创建表失败",err)
		return
	}*//*
//conn.Exec("insert into user (name,password) values (?,?)","tom","123")

res,err:=conn.Query("select name from user")
var name string
for res.Next(){
	res.Scan(&name)
	beego.Info(name)
}*/

orm.RegisterDataBase("default","mysql","root:123456@(127.0.0.1:3306)/beegotest?charset=utf8")

//后台new表 用户表 和 文章表
orm.RegisterModel(new(Test),new(Article),new(ArticleType))


orm.RunSyncdb("default",false,true)

}