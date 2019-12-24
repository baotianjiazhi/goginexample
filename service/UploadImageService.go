package service

import (
	"ginexample/pkg/upload"
	"ginexample/serializer"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type UploadImageService struct {
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}


func (service *UploadImageService)Upload(c *gin.Context) serializer.Response {
	file, err := service.Avatar.Open()
	if err != nil {
		return serializer.Response{
			Status: 400001,
			Msg: "文件上传失败",
			Error: err.Error(),
		}
	}
	imageName := upload.GetImageName(service.Avatar.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		return serializer.Response{
			Status: 40001,
			Msg:    "文件过大且后缀只能为jpg或png",
		}
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		return serializer.Response{
			Status: 400001,
			Msg: "文件上传失败",
			Error: err.Error(),
		}
	}

	if err := c.SaveUploadedFile(service.Avatar, src); err != nil {
		return serializer.Response{
			Status: 400001,
			Msg: "文件上传失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Msg: "文件上传成功",
		Data: map[string]string{
			"image_url":      upload.GetImageFullUrl(imageName),
			"image_save_url": savePath + imageName,
		},
	}
}