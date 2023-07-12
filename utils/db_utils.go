package utils

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

var (
	DB *sql.DB
)

func InitDB() error {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		return errors.New("env MYSQL_DSN must set")
	}
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}
	return nil
}

var Db1 *gorm.DB
var mysqlLogger logger.Interface

func init() {
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	DbName := "bookstore"
	timeOut := "10s"

	mysqlLogger = logger.Default.LogMode(logger.Info)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		username, password, host, port, DbName, timeOut)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		//SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			//	TablePrefix:   "",   //表名前缀
			SingularTable: true, //是否为单数表名
			//	NoLowerCase:   true, //不要小写转换
		},
		//Logger: mysqlLogger,
	})
	if err != nil {
		panic("连接数据库失败,err" + err.Error())

	}
	//连接成功
	Db1 = db
	Db1 = db.Session(&gorm.Session{
		Logger: mysqlLogger,
	})
}
