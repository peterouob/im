package models

import "gorm.io/gorm"

//群信息
type GroupBasic struct {
	gorm.Model
	Name    string //群名稱
	OwnerId uint   //群屬於誰
	Icon    string //群圖片
	Type    int
	Desc    string //描述
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
