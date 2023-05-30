package mysql

import (
	"NetworkDisk/config"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {
	var err error
	mysqlConfig := config.GlobalConfig.Databases.Mysql
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf("%v:%v@tcp(%v%v)/%v?charset=%v&parseTime=True&loc=Local",
		mysqlConfig.Account, mysqlConfig.Password, mysqlConfig.URL, mysqlConfig.Port,
		mysqlConfig.DbName, mysqlConfig.Charset)

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // 慢 SQL 阈值
	// 		LogLevel:                  logger.Info, // 日志级别
	// 		IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
	// 		Colorful:                  false,       // 禁用彩色打印
	// 	},
	// )

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger:                 newLogger,
		SkipDefaultTransaction: true, // 禁用默认事务
		PrepareStmt:            true, // 缓存预编译语句
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   true, // skip the snake_casing of names
			SingularTable: true,
		},
	})
	if err != nil {
		log.Println("数据库连接失败")
		log.Println(err)
		os.Exit(1)
	}

	log.Println("数据库连接成功")

	//根据*grom.DB对象获得*sql.DB的通用数据库接口
	sqlDb, _ := DB.DB()
	sqlDb.SetMaxIdleConns(mysqlConfig.MaxIdleConns) //设置最大连接数
	sqlDb.SetMaxOpenConns(mysqlConfig.MaxOpenConns) //设置最大的空闲连接数
}
