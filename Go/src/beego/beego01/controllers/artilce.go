package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	c.Ctx.WriteString("新闻列表") //直接向页面返回内容
}
func (c *ArticleController) AddArticle() { //方法名首字母必须大写
	c.Ctx.WriteString("增加新闻")
}
func (c *ArticleController) EditArticle() {
	//获取get传值
	// id := c.GetString("id")
	// Fmt.Printf("值：%v,类型：%T", id, id)
	// beego.Info(id)
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		c.Ctx.WriteString("传入参数错误")
		return
	}
	fmt.Printf("值：%v,类型：%T\n", id, id)
	c.Ctx.WriteString("编辑新闻")
}
