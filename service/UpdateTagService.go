package service

import (
	"ginexample/models"
	"ginexample/serializer"
)

type UpdateTagService struct {
	Name string `form:"name" json:"name" binding:"required,min=2,max=30"`
	ModifiedBy string `form:"modified_by" json:"modified_by" binding:"required,min=5,max=30"`
}


// 更新文章
func (service *UpdateTagService) Update(id string) serializer.Response {
	var tag models.Tag
	err := models.DB.First(&tag, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg: "标签不存在",
			Error: err.Error(),
		}
	}

	var tag_copy models.Tag
	models.DB.Where("name = ?", service.Name).First(&tag_copy)
	if  tag_copy.ID > 1{
		return serializer.Response{
			Status: 400002,
			Msg: "标签名称已存在",
		}
	}

	_, err = models.GetUserByName(service.ModifiedBy)
	if err != nil {
		return serializer.Response{
			Status: 40002,
			Msg: "查询不到创建用户",
			Error: err.Error(),
		}
	}

	tag.Name = service.Name
	tag.ModifiedBy = service.ModifiedBy

	err = models.DB.Save(&tag).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg: "标签保存失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildTag(tag),
	}
}