package service

import (
	"ginexample/models"
	"ginexample/serializer"
)

type ShowArticlesService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

func (service *ShowArticlesService) Show() serializer.Response {
	articles := []models.Article{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 10
	}

	if err := models.DB.Model(&models.Article{}).Count(&total).Error; err != nil{
		return serializer.Response{
			Status: 50000,
			Msg: "数据库连接错误",
			Error: err.Error(),
		}
	}

	if err := models.DB.Limit(service.Limit).Offset(service.Start).Find(&articles).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg: "数据库连接错误",
			Error: err.Error(),
		}
	}


	return serializer.BuildListResponse(serializer.BuildArticles(articles), uint(total))
}
