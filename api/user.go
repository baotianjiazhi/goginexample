package api

import (
	"fmt"
	"ginexample/serializer"
	"ginexample/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

func CreateUser(c *gin.Context) {
	var service service.CreateUserService
	if err := c.ShouldBind(&service); err == nil {
		if user, err := service.Register(); err != nil {
			c.JSON(200, err)
		} else {
			res := serializer.BuildUserResponse(user)
			c.JSON(200, res)
		}
	} else {
		fmt.Println(err)
		c.JSON(200, err)
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		if user, err := service.Login(); err != nil {
			c.JSON(200, err)
		} else {
			// 设置Session
			s := sessions.Default(c)
			s.Clear()
			s.Set("user_id", user.ID)
			s.Save()

			res := serializer.BuildUserResponse(user)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


// 查询当前登陆用户信息
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUser(*user)
	c.JSON(200, res)
}

// 退出用户
func Logout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "登出成功",
	})
}