package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
)

// 默认数据库实例
var defaultDB *gorm.DB

// 初始化默认数据库
func InitDefaultDB(source, driver string, cfg *gorm.Config) (db *gorm.DB, err error) {
	if cfg == nil {
		//cfg不能为nil
		cfg = &gorm.Config{}
	}
	db, err = NewDB(source, driver, cfg)
	if err != nil {
		return nil, err
	}
	defaultDB = db
	fmt.Println("db init success")
	return defaultDB, nil
}

// 新建数据库
func NewDB(source, driver string, cfg *gorm.Config) (db *gorm.DB, err error) {
	d := strings.ToLower(driver)
	switch d {
	case "mysql":
		db, err = initMysqlDB(source, cfg)
	case "postgres":
		db, err = initPostgresDB(source, cfg)
	case "sqlite":
		db, err = initSqliteDB(source, cfg)
	}
	return db, err
}

// 初始化默认mysql数据库
func initMysqlDB(source string, cfg *gorm.Config) (db *gorm.DB, err error) {
	d, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	return gorm.Open(mysql.New(mysql.Config{Conn: d}), cfg)
}

// 初始化默认mysql数据库
func initPostgresDB(source string, cfg *gorm.Config) (db *gorm.DB, err error) {
	d, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	return gorm.Open(postgres.New(postgres.Config{Conn: d}), cfg)
}

// 初始化默认sqlite数据库
func initSqliteDB(dbfile string, cfg *gorm.Config) (db *gorm.DB, err error) {
	return gorm.Open(sqlite.Open(dbfile), cfg)
}

// 关闭数据库。
func Close() (err error) {
	if defaultDB != nil {
		//return defaultDB.Close()
	}
	return nil
}

// 获取数据库连接对象。
func GetDB() *gorm.DB {
	return defaultDB.Debug()
}
