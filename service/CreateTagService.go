package service

import (
	"ginexample/models"
	"ginexample/serializer"
	"log"
)

type CreateTagService struct {
	Name      string `form:"name" json:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" json:"created_by" binding:"required,min=5,max=30"`
}


func (service *CreateTagService) CreateTag() serializer.Response {
	tag := models.Tag{
		Name: service.Name,
		CreatedBy: service.CreatedBy,
		Status: models.Active,
	}

	_, err := models.GetUserByName(service.CreatedBy)
	if err != nil {
		return serializer.Response{
			Status: 40002,
			Msg: "查询不到创建用户",
			Error: err.Error(),
		}
	}

	exists, err := models.ExistByName(service.Name)
	log.Println(exists, err)
	if exists == true {
		return serializer.Response{
			Status: 40002,
			Msg: "tag已存在",
			Error: err.Error(),
		}
	}

	if err := models.DB.Create(&tag).Error; err != nil {
		return serializer.Response{
			Status: 40002,
			Msg: "创建失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildTag(tag),
	}

}