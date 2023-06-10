package controllers

import (
	"beego02/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	/*//创建orm对象
	o := orm.NewOrm()
	//插入结构体对象
	user := models.User{}
	//对结构体对象赋值
	user.Name = "1234565"
	user.Pwd = "123456"
	//插入
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入失败", err)
		return
	}
	c.Ctx.WriteString("创建数据成功")*/

	////有orm对象
	//o := orm.NewOrm()
	////查询的对象
	//user := models.User{}
	//user.Name="111"
	////指定查询对象字段值
	///*user.Id = 1
	////查询
	//err := o.Read(&user)*/
	//err := o.Read(&user, "Name")
	//if err != nil {
	//	beego.Info("查询失败", err)
	//	return
	//}
	//beego.Info("查询成功", user)
	//c.Ctx.WriteString("查询数据成功")

	//o := orm.NewOrm()
	//user := models.User{}
	//user.Id = 1
	//_, err := o.Delete(&user)
	//if err != nil {
	//	beego.Info("删除错误")
	//	return
	//}
	//c.Ctx.WriteString("删除数据成功")

	/*o := orm.NewOrm()
	querySet := o.QueryTable("user")
	result := make([]models.User, 0, 10)
	_, err := querySet.All(&result)
	if err != nil {
		beego.Info("查询出错啦:", err)
	}
	for index, user := range result {
		c.Ctx.WriteString(Fmt.Sprintf("第%d条是%v\n", index, user))
	}*/

	//高级查询
	//o := orm.NewOrm()
	//ids := []int{control, 3, memory}
	//res, err := o.Raw("SELECT name FROM user WHERE id IN (?, ?, ?)", ids).Exec()
	//if err == nil {
	//	num, _ := res.RowsAffected()
	//	c.Ctx.WriteString(Fmt.Sprintf("Affect nums:", num))
	//}

	o := orm.NewOrm()
	user := models.User{}
	r := o.Raw("SELECT id,name,pwd FROM user WHERE id = ?", 3)
	r.QueryRow(&user)
	c.Ctx.WriteString(fmt.Sprintf("user is %v", user))

	//o := orm.NewOrm()
	//var users []models.User
	//num, err := o.Raw("SELECT id, user_name FROM user WHERE id = ?", 1).QueryRows(&users)
	//if err == nil {
	//	c.Ctx.WriteString(Fmt.Sprintf("user nums:%v ", num))
	//}

}
