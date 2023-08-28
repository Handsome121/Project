package main

import (
	"fmt"
	"gin-chat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123456@tcp(127.0.0.1:3306)/user?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                          // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                         // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                         // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                         // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                        // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// ----------------------------数据库连接池----------------------------
	sqlDB, err := db.DB()
	// SetMaxIdleCons 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenCons 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	fmt.Println("success to link mysql")

	err = db.AutoMigrate(&models.UserBasic{})
	if err != nil {
		panic("auto migrate fail")
	}
}
