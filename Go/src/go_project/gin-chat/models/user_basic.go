package models

import (
	"fmt"
	"gin-chat/initial"
	"gin-chat/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string `valid:"matches(^13\\d{9}$)"`
	Email         string `valid:"email"`
	Identify      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartBeatTime time.Time
	LogOutTime    time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogOut      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	var users []*UserBasic

	initial.DB.Find(&users)
	for _, v := range users {
		fmt.Println(v)
	}
	return users

}

func CreateUser(user UserBasic) *gorm.DB {

	return initial.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {

	return initial.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {

	return initial.DB.Model(&user).Updates(UserBasic{Name: user.Name, Password: user.Password})
}

func FindUserByName(name string) UserBasic {
	var user UserBasic
	initial.DB.Where("name = ?", name).First(&user)

	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	initial.DB.Model(&user).Where("id = ?", user.ID).Update("identify", temp)

	return user
}

func FindUserByNameAndPwd(name, passWord string) UserBasic {
	var user UserBasic
	initial.DB.Where("name = ? and password = ?", name, passWord).First(&user)

	return user
}

func FindUserByPhone(phone string) UserBasic {
	var user UserBasic
	initial.DB.Where("phone = ?", phone).First(&user)

	return user
}

func FindUserByEmail(email string) UserBasic {
	var user UserBasic
	initial.DB.Where("email = ?", email).First(&user)

	return user
}
