package service

import (
	"ginexample/models"
	"ginexample/serializer"
)

type CreateArticleService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=3000"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,max=255"`
}

func (service *CreateArticleService) Create() serializer.Response {
	article := models.Article{
		Title:   service.Title,
		Content: service.Content,
		Status:  models.Active,
		CoverImageUrl: service.CoverImageUrl,
	}

	if err := models.DB.Create(&article).Error; err != nil {
		return serializer.Response{
			Status: 40002,
			Msg:    "上传失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildArticle(article),
	}
}
