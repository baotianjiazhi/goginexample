package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName string
	PassowordDigest string
	Nickname string
	Status string
}


const(
	// Passowrdcost 密码加密难度
	PasswordCost = 12
	// Suspend 被封禁账户
	Suspend string = "suspend"
)

// 通过ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// 通过Name获取用户
func GetUserByName(UserName interface{}) (User, error) {
	var user User
	result := DB.Where("user_name = ?", UserName).First(&user)
	return user, result.Error
}

// 设置密码
func (user *User)SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}

	user.PassowordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User)CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassowordDigest), []byte(password))
	return err == nil
}