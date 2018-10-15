package main

import (
	_"beego1/routers"
	"github.com/astaxie/beego"
	_"beego1/models"
)

func main() {
	beego.AddFuncMap("prepage",ShowPrePage)
	beego.AddFuncMap("nextpage",ShowNextPage)
	beego.Run()
}
func ShowPrePage(pageIndex int) int {
	if pageIndex==1{
		return pageIndex
	}
		return pageIndex-1
}
func ShowNextPage(pageIndex,pagecount int) int {
	if pageIndex==pagecount{
		return pageIndex
	}
	return pageIndex+1
}