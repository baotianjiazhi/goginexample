package service

import (
	"ginexample/models"
	"ginexample/serializer"
)

type DeleteArticleService struct {
}


func (service DeleteArticleService) Delete(id string) serializer.Response {
	var article models.Article
	err := models.DB.First(&article, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg: "文章不存在",
			Error: err.Error(),
		}
	}

	article.Status = models.Inactive
	err = models.DB.Delete(&article).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg: "文章删除失败",
			Error: err.Error(),
		}
	}


	return serializer.Response{}
}