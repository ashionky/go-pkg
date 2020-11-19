/**
 * @Author pibing
 * @create 2020/11/19 3:27 PM
 */

package model

import "go-pkg/pkg/db"

func InitTables()  {
	db.GetDB().AutoMigrate(User{})
}
