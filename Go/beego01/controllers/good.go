package controllers

import (
	"encoding/xml"
	"fmt"

	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {
	c.Data["title"] = "你好beego" //绑定数据
	c.Data["num"] = 12          //绑定数据
	c.TplName = "goods.tpl"
}
func (c *GoodsController) DoAdd() { //post
	c.Ctx.WriteString("执行增加操作")
}

type Product struct {
	Title   string `form:"title" xml:"title"`
	Content string `form:"content" xml:"content"`
}

func (c *GoodsController) DoEdit() { //put
	// title := c.GetString("title")
	p := Product{}
	if err := c.ParseForm(&p); err != nil {
		c.Ctx.WriteString("put提交失败")
		return
	}
	fmt.Printf("%#v", p)
	c.Ctx.WriteString("执行修改操作")
}
func (c *GoodsController) DoDelete() { //delete
	c.Ctx.WriteString("执行删除操作")
}
func (c *GoodsController) Xml() {
	p := Product{}
	// str := string(c.Ctx.Input.RequestBody)
	// beego.Info(str)
	// p := Product{}
	// c.Ctx.WriteString(str)
	var err error
	if e := xml.Unmarshal(c.Ctx.Input.RequestBody, &p); e != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		c.Data["json"] = p
		c.ServeJSON()
	}

}
