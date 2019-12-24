package service

import (
	"ginexample/models"
	"ginexample/serializer"
)

type UpdateArticleService struct {
	Title string `form:"title" json:"title" binding:"min=2,max=30"`
	Content string `form:"content" json:"content" binding:"required,max=3000"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,max=255"`
}


// 更新文章
func (service *UpdateArticleService) Update(id string) serializer.Response {
	var article models.Article
	err := models.DB.First(&article, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg: "文章不存在",
			Error: err.Error(),
		}
	}

	if service.Title != "" {
		article.Title = service.Title
	}

	if service.CoverImageUrl != "" {
		article.CoverImageUrl = service.CoverImageUrl
	}
	article.Content = service.Content
	err = models.DB.Save(&article).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg: "视频保存失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildArticle(article),
	}
}