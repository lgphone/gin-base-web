package mysqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
	"younghe/config"
)

var DB *xorm.Engine

func Setup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?timeout=10s&charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Mysql.User,
		config.Config.Mysql.Password,
		config.Config.Mysql.Host,
		config.Config.Mysql.Port,
		config.Config.Mysql.DB,
	)
	DB, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// 设置统一表前缀
	mapper := names.NewPrefixMapper(names.GonicMapper{}, config.Config.Mysql.TablePrefix)
	DB.SetTableMapper(mapper)
	// 设置表字段映射为蛇形命名
	DB.SetColumnMapper(names.GonicMapper{})

	DB.ShowSQL(true) // 设置显示sql语句  会在控制台当中展示
	logFile, err := os.Create(config.Config.SqlLogFilePath)
	if err != nil {
		panic(err)
	}
	DB.SetLogger(log.NewSimpleLogger(logFile)) // 设置日志级别
	DB.SetLogLevel(log.LOG_DEBUG)
	DB.ShowSQL(true)
	DB.SetMaxOpenConns(500) // 设置最大链接数量
	DB.SetMaxIdleConns(100) // 连接池最大空闲连接数量
	// check connection
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}
