package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"time"
	"github.com/astaxie/beego/orm"
	"beego1/models"
	"math"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) ShowArticlelist() {


	userName:=this.GetSession("username")
	if userName==nil{
		this.Redirect("/login",302)
		return
	}

	//获取数据
	o:=orm.NewOrm()

	qs:=o.QueryTable("Article")

	var articles []models.Article

	typeName:=this.GetString("select")

	var count int64

	pageSize:=2

	//获取页码
	pageIndex,err:=this.GetInt("pageIndex")

	if err!=nil{
		pageIndex=1
	}
	start:=(pageIndex-1)*pageSize

	if typeName==""{
		count,_=qs.Count()
	}else {
		count,_=qs.Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).Count()
	}
	pageCount:=math.Ceil( float64(count)/float64(pageSize))

	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)
	this.Data["types"]=types
	beego.Info(typeName)
	if typeName=="" {
		qs.Limit(pageSize,start).All(&articles)
	}else {
		qs.Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articles)
	}
    this.Data["typeName"]=typeName
	this.Data["pageIndex"]=pageIndex
	//末页
	this.Data["pageCount"]=int(pageCount)
	this.Data["count"]=count
	//传递显示每行文章
	this.Data["articles"]=articles

	userlay:=this.GetSession("username")
	this.Data["username"]=userlay.(string)
	this.Layout="layout.html"
	this.TplName="index.html"
}

func (this *ArticleController) ShowaddArticle() {
	o:=orm.NewOrm()

	var types []models.ArticleType

	o.QueryTable("ArticleType").All(&types)
	this.Data["types"]=types

	userlay:=this.GetSession("username")
	this.Data["username"]=userlay.(string)
	this.Layout="layout.html"
	this.TplName="add.html"
}

func (this *ArticleController) HandleaddArticle() {

	articleName:=this.GetString("articleName")

	content:=this.GetString("content")

	if articleName==""||content==""{
		this.Data["errmsg"]="添加数据不完整"
		this.TplName="add.html"
		return
	}

   filepath:=UploadFile(&this.Controller,"uploadname")

	o:=orm.NewOrm()

	var article models.Article

	article.ArtiName=articleName
	article.Acontent=content
	article.Aimg=filepath

	typeNAme:=this.GetString("select")

	var articleType models.ArticleType

	articleType.TypeName=typeNAme

	o.Read(&articleType,"TypeName")

	article.ArticleType=&articleType

	o.Insert(&article)

	this.Redirect("/article/showArticlelist",302)
}

func (this *ArticleController) ShowArticleDetail(){


	id,err:=this.GetInt("articleId")

	if err !=nil{
		beego.Info("传递的链接错误")
	}
	o:=orm.NewOrm()
	var article models.Article
	article.Id=id
	o.QueryTable("Article").RelatedSel("ArticleType").Filter("Id",id).One(&article)

    article.Acount+=1
    o.Update(&article)

    //多对多插入
    m2m:=o.QueryM2M(&article,"User")

	userName:=this.GetSession("username")

	if userName==nil{
		this.Redirect("/login",302)
		return
	}
	var user models.Test
	user.Name=userName.(string)

	o.Read(&user,"Name")
	m2m.Add(user)

	//o.LoadRelated(&article,"User")
	var users []models.Test
	o.QueryTable("Test").Filter("Articles__Article__Id",id).Distinct().All(&users)
	this.Data["users"]=users
	this.Data["article"]=article

	userlay:=this.GetSession("username")
	this.Data["username"]=userlay.(string)
	this.Layout="layout.html"
	this.TplName="content.html"
}

func (this *ArticleController) ShowUpdateArticle() {
	//获取数据
	id,err:=this.GetInt("articleId")
	//检验数据
	if err!=nil{
		beego.Info("传递的链接不正确")
	}

	//数据处理
	o:=orm.NewOrm()
	var article models.Article
	article.Id=id
	o.Read(&article)

	//返回视图

	this.Data["article"]=article

	userlay:=this.GetSession("username")
	this.Data["username"]=userlay.(string)
	this.Layout="layout.html"
	this.TplName="update.html"



}

//封装上传文件方式
func UploadFile(this *beego.Controller,filepath string) string {
	file,head,err:=this.GetFile(filepath)
	if head.Filename==""{
		return "NOIMG"
	}
	if err!=nil{
		this.Data["errmsg"]="文件上传失败"
		this.TplName="add.html"
		return ""
	}
	defer  file.Close()

	if head.Size>50000000000{
		this.Data["errmsg"]="文件太大"
		this.TplName="add.html"
		return ""
	}
	ext:=path.Ext(head.Filename)
	if ext!=".JPG" && ext!=",png" && ext!=".jpeg" &&ext!=".jpg"&&ext!=".PNG"{
		this.Data["errmsg"]="文件格式不正确"
		this.TplName="add.html"
		return ""

		}

	filename:=time.Now().Format("2006-01-02-15:04:05")+ext

	this.SaveToFile(filepath,"./static/img/"+filename)
     return  "/static/img/"+filename
}
//处理编辑界面数据

func (this *ArticleController) HandleUpdateArticle(){
	//获取数据
	id,err:=this.GetInt("articleId")
	articleName:=this.GetString("articleName")

	content:=this.GetString("content")

	filepath:=UploadFile(&this.Controller,"uploadname")
	//数据校验
	if articleName=="" || err!=nil || content=="" || filepath==""{
		beego.Info("请求错误")
		return
	}

	o:=orm.NewOrm()
	var article models.Article
	article.Id=id

	err=o.Read(&article)
	if err!=nil{
		beego.Info("更新的文章不存在")
		return
	}

	article.ArtiName=articleName
	article.Acontent=content
	if filepath!="NOIMG"{
		article.Aimg=filepath
	}



	o.Update(&article)

	//返回视图
	this.Redirect("/article/showArticlelist",302)

}

func (this *ArticleController)  DeleteArticle(){
	//获取数据
	id,err:=this.GetInt("articleId")
	//校验数据
	if err!=nil{
		beego.Info("删除文章错误")
		return
	}
	//数据处理
	o:=orm.NewOrm()
	var article models.Article
	article.Id=id
	o.Delete(&article)

	//返回视图
	this.Redirect("/article/showArticlelist",302)
}

func (this *ArticleController) ShowAddType() {

	o:=orm.NewOrm()
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)

	this.Data["types"]=types
	userlay:=this.GetSession("username")
	this.Data["username"]=userlay.(string)
	this.Layout="layout.html"
	this.TplName="addType.html"
}

func (this *ArticleController) HandleAddType() {
	typeName:=this.GetString("typeName")

	if typeName==""{
		beego.Info("获取信息不完整")
		return
	}

	o:=orm.NewOrm()
	var article models.ArticleType
	article.TypeName=typeName
	o.Insert(&article)

	this.Redirect("/article/addType",302)
}

func (this *ArticleController) DeleteType() {
	id,err:=this.GetInt("id")
	if err!=nil{
		beego.Error("删除类型错误")
		return
	}

	o:=orm.NewOrm()
	var articleType models.ArticleType
	articleType.Id=id

	o.Delete(&articleType)

	this.Redirect("/article/addType",302)
}