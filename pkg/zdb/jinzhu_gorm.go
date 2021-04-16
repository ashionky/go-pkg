package zdb

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strings"
)

// 默认数据库实例
var defaultDB *gorm.DB

// 初始化默认数据库
func InitDefaultDB(source, driver string) (db *gorm.DB, err error) {
	db, err = NewDB(source, driver)
	if err != nil {
		return nil, err
	}
	defaultDB = db
	defaultDB.LogMode(true) //打印详细sql日志
	fmt.Println("db init success")
	return defaultDB, nil
}

// 新建数据库
func NewDB(source, driver string) (db *gorm.DB, err error) {
	d := strings.ToLower(driver)
	db, err = gorm.Open(d, source)
	return db, err
}

// 关闭数据库。
func Close() (err error) {
	if defaultDB != nil {
		return defaultDB.Close()
	}
	return nil
}

// 获取数据库连接对象。
func GetDB() *gorm.DB {
	return defaultDB.Debug()
}
