package initial

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB  *gorm.DB
	V   *viper.Viper
	Red *redis.Client
)

func InitConfig() {
	V = viper.New()

	V.SetConfigName("app")
	V.SetConfigType("yml")
	V.AddConfigPath("E:\\Project\\Go\\src\\go_project\\gin-chat\\config")

	err := V.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func InitMysql() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			LogLevel:                  logger.Info,
		},
	)

	dsn := V.GetString("mysql.dns")
	DB, _ = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: newLogger,
	})

	// ----------------------------数据库连接池----------------------------
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("sql DB fail")
	}
	// SetMaxIdleCons 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenCons 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	fmt.Println("success to link mysql")
}

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         V.GetString("redis.addr"),
		Password:     V.GetString("redis.password"),
		DB:           V.GetInt("redis.DB"),
		PoolSize:     V.GetInt("redis.poolSize"),
		MinIdleConns: V.GetInt("redis.minIdleConn"),
	})

	pong, err := Red.Ping().Result()
	if err != nil {
		fmt.Println("init redis connect failed......", err)
	}
	fmt.Println("redis init success......", pong)

}
