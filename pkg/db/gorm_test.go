/**
 * @Author pibing
 * @create 2020/11/16 9:28 AM
 */

package db

import (
	"fmt"
	"go-pkg/pkg/cfg"
	"gorm.io/gorm"
	"testing"
)

type Test struct {
	gorm.Model
	Name string `json:"name"`
}

//自定义表名
func (Test) TableName() string {
	return "test_b"
}

func TestInitDefaultDB(t *testing.T) {
	var configFile = "../../conf/dev.yml"
	_ = cfg.Initcfg(configFile)
	var config = cfg.GetConfig()
	driver := "mysql"

	user := config.Mysql.User
	password := config.Mysql.Password
	host := config.Mysql.Host
	dbname := config.Mysql.Dbname
	charset := config.Mysql.Charset
	sqlconnStr := fmt.Sprintf("%v:%v@(%v)/%v?charset=%v&parseTime=True&loc=Local",
		user, password, host, dbname, charset)
	_, _ = InitDefaultDB(sqlconnStr, driver, nil)
	defaultDB.AutoMigrate(Test{})
	tt := Test{
		Name: "3323231d",
	}
	defaultDB.Table("test_b").Save(&tt)
	list := []Test{}
	defaultDB.Table("test_b").Find(&list)

	fmt.Print("id===", list)

}
