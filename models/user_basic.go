package models

import (
	"fmt"
	"gin-chat/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string    `json:"name"`
	Password      string    `json:"password"`
	Phone         string    `json:"phone"` //\\d=>[0-9]
	Email         string    `valid:"email" json:"email"`
	Identity      string    `json:"identity"`  // 唯一標示
	ClientIp      string    `json:"client_ip"` //設備
	ClientPort    string    `json:"client_port"`
	Salt          string    //加密
	LoginTime     time.Time `json:"login_time"`
	HeartbeatTime time.Time `json:"heartbeat_time"` //心跳時間
	LogOutTime    time.Time `json:"log_out_time" gorm:"column:login_out_time"`
	IsLogOut      bool      `json:"is_log_out"`
	DeviceInfo    string    `json:"device_info"` //設備信息
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func FindUserByNameAndPwd(name, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name= ? and password=? ", name, password).First(&user)
	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.Md5Encode(str)
	utils.DB.Model(&user).Where("id=?", user.ID).Update("Identity", temp)
	return user
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name= ?", name).First(&user)
	return user
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(&user)
}
