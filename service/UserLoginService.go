package service

import (
	"ginexample/models"
	"ginexample/serializer"
)

type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 用户登录函数
func (service *UserLoginService) Login() (models.User, *serializer.Response) {
	var user models.User

	if err := models.DB.Where("user_name = ? ", service.UserName).First(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: 400001,
			Msg: "账号或密码错误",
		}
	}

	if user.CheckPassword(service.Password) == false {
		return user, &serializer.Response{
			Status: 40001,
			Msg: "账号或密码错误",
		}
	}

	return user, nil
}