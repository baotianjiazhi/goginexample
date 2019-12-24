package service

import (
	"ginexample/models"
	"ginexample/serializer"
)

type DeleteTagService struct {
}


func (service *DeleteTagService) DeleteTag(id string) serializer.Response {
	var tag models.Tag
	err := models.DB.First(&tag, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg: "标签不存在",
			Error: err.Error(),
		}
	}

	tag.Status = models.Inactive
	err = models.DB.Save(&tag).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg: "标签删除失败",
			Error: err.Error(),
		}
	}

	err = models.DB.Delete(&tag).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg: "标签删除失败",
			Error: err.Error(),
		}
	}


	return serializer.Response{
		Msg: "删除成功",
	}
}