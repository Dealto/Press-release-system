package routers

import (
	"beego1/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	beego.InsertFilter("/article/*",beego.BeforeExec,Filter)
    beego.Router("/", &controllers.MainController{},"get:Get;post:ShowGet")

    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandlePost")

	beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HandleLogin")

    beego.Router("/article/showArticlelist",&controllers.ArticleController{},"get:ShowArticlelist")

	beego.Router("/article/addArticle",&controllers.ArticleController{},"get:ShowaddArticle;post:HandleaddArticle")

	beego.Router("/article/showArticleDetail",&controllers.ArticleController{},"get:ShowArticleDetail")

	beego.Router("/article/updateArticle",&controllers.ArticleController{},"get:ShowUpdateArticle;post:HandleUpdateArticle")

	beego.Router("/article/deleteArticle",&controllers.ArticleController{},"get:DeleteArticle")

	beego.Router("/article/addType",&controllers.ArticleController{},"get:ShowAddType;post:HandleAddType")

	beego.Router("/Logout",&controllers.UserController{},"get:LogOut")


	beego.Router("/article/deleteType",&controllers.ArticleController{},"get:DeleteType")

	}

var Filter = func(ctx *context.Context) {

	username:=ctx.Input.Session("username")
	if username==nil{
		ctx.Redirect(302,"/login")
		return
	}
}