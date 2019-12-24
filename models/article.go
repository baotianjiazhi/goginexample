package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title         string `json:"title"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	Status        string `json:"status"`
}

func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := DB.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}
