package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("api接口" + id)

}
