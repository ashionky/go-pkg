/**
 * @Author pibing
 * @create 2020/11/14 1:24 PM
 */

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	// 手机号
	Phone string `json:"phone"`
	// 用户名
	Name string `json:"name"`
	// 密码
	Password string `json:"password"`
	// 用户状态，1为正常，2为禁用
	State uint `json:"state"`
}

//自定义表名
func (User) TableName() string {
	return "t_user"
}
