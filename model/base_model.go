/**
 * @Author pibing
 * @create 2020/11/5 4:40 PM
 */

package model

import (
	"gorm.io/gorm"
	"time"
)

//可用ModelBase 替换gorm.Model  设置json格式/字段
type ModelBase struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
