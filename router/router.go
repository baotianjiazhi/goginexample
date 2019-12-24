package router

import (
	"ginexample/api"
	"ginexample/middleware"
	"ginexample/pkg/upload"
	"github.com/gin-gonic/gin"
	"ginexample/setting"
	"net/http"
)

func Newrouter() *gin.Engine {
	r := gin.Default()

	// 初始化session
	r.Use(middleware.Session(setting.AppSetting.SessionSecret))
	// 获取当前登陆用户信息
	r.Use(middleware.CurrentUser())
	// 让前端可以获取图片
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	// 生成二维码

	//r.LoadHTMLGlob("templates/**/*")
	v1 := r.Group("/api/v1")
	{
		// 测试
		v1.GET("/ping", api.Ping)

		// 微博标签
		v1.POST("/tag", api.CreateTag)
		// 删除标签
		v1.DELETE("/tag/:id", api.DeleteTag)
		// 更新标签
		v1.PUT("/tag/:id", api.UpdateTag)

		// 上传微博内容
		v1.POST("/article", api.CreateArticle)
		// 删除微博内容
		v1.DELETE("/article/:id", api.DeleteArticle)
		// 查询单个微博内容
		v1.GET("/article/:id", api.ListArticle)
		// 查询所有微博
		v1.GET("/articles", api.ShowArticles)
		// 修改微博内容
		v1.PUT("/article/:id", api.UpdateArticle)
		//  上传文件
		v1.POST("/upload", api.UploadImage)

		// 用户注册
		v1.POST("/user/register", api.CreateUser)
		// 用户登陆
		v1.POST("/user/login", api.UserLogin)

		// 生成二维码
		v1.POST("/articles/poster/generate", api.GenerateArticlePoster)

		// 需要登陆才能进行
		authed := v1.Group("/")
		authed.Use(middleware.AuthRequired())
		{
			// 查询个人信息
			authed.GET("/user/me", api.UserMe)
			// 登出
			authed.DELETE("/user/logout", api.Logout)
		}
	}

	return r
}
