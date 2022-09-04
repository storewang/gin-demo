package utils

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/shitou/go-demo-gin/web/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbUtil struct {
	Db        *gorm.DB
	dbhandler *sql.DB
}

func (d *DbUtil) Init(conf *config.DbConf) {
	dsn := fmt.Sprint(conf.User, ":", conf.Passwd, "@tcp(", conf.Host, ")/", conf.Schema, "?charset=utf8mb4&parseTime=True&loc=Local")
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败", err.Error())
	}

	// 连接池相关设置
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("数据库打开失败", err.Error())
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	d.Db = db
	d.dbhandler = sqlDB
}

func (d *DbUtil) Close() {
	d.dbhandler.Close()
}
