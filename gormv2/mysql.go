package gormv2

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	db *gorm.DB
)

// InitMysql [账号]:[密码]@tcp([地址]:[端口]) db.table([库名].[表名])
func InitMysql(dbUrl string) {
	url := fmt.Sprintf("%v/?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", dbUrl)
	db = NewMysql(url, 10)
}

func GetMainDB() *gorm.DB {
	return db
}

func NewMysql(args string, maxCon int) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(fmt.Sprintf("Got errors when connect database, the errors is '%v'", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	idle := maxCon
	if maxCon/3 > 10 {
		idle = maxCon / 3
	}
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(idle)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(maxCon)

	return db
}
