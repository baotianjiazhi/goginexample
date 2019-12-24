package service

import (
	"ginexample/models"
	"ginexample/serializer"
)

// ListArticleService 视频列表服务
type ListArticleService struct {
}


// List 内容列表
func (service *ListArticleService) List(id string) serializer.Response {
	article := models.Article{}
	err := models.DB.First(&article, id).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg: "没有查询到对应的内容",
			Error: err.Error(),
		}
	}


	return serializer.Response{
		Data: serializer.BuildArticle(article),
		Msg: "成功查找到内容",
	}
}