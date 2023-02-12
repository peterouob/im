package models

import "gorm.io/gorm"

//消息
type Message struct {
	gorm.Model
	FromId   uint   //發送者
	TargetId uint   //接收者
	Type     string //消息類席 群聊 私聊 廣播
	Media    string //消息類型 文字 圖片 音頻
	Content  string //消息內容
	Pic      string
	Url      string
	Desc     string //描述
	Amount   int    //其他數字統計
}

func (table *Message) TableName() string {
	return "message"
}
