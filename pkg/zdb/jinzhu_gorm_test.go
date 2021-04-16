/**
 * @Author pibing
 * @create 2020/12/18 4:43 PM
 */

package zdb

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-pkg/pkg/cfg"
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
	_, _ = InitDefaultDB(sqlconnStr, driver)
	defaultDB.AutoMigrate(Test{})
	//tt:=Test{
	//	Name:  "33",
	//}
	count := 0
	defaultDB.Table("test_b").Count(&count)
	fmt.Println("1===", count)
	//defaultDB.Table("test_b").Save(&tt)

	defaultDB.Table("test_b").Where("id=?", 3).Delete(&Test{})

	defaultDB.Table("test_b").Count(&count)

	fmt.Println("2===", count)

}
