package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //驱动引入
)

//表的设计
type User struct {
	Id   int
	Name string
	Pwd  string
}

func init() {
	//注册驱动
	//orm.RegisterDriver("sqlite3", orm.DRSqlite)
	//设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	//映射model数据
	orm.RegisterModel(new(User))
	//生成表
	orm.RunSyncdb("default", false, true)
}
