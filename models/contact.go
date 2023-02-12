package models

import "gorm.io/gorm"

//人緣關係
type Contact struct {
	gorm.Model
	OwnerId  uint   //誰的關係信息
	TargetId uint   //對應誰
	Type     int    //對應的類型 好友 其他
	Desc     string //描述訊息
}

func (table *Contact) TableName() string {
	return "contact"
}
