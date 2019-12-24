package api

import (
	"ginexample/pkg/qrcode"
	"ginexample/serializer"
	"ginexample/service"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
)

const (
	QRCODE_URL = "https://baotianjiazhi.github.io/"
)

// CreateArticle 创建微博内容
func CreateArticle(c *gin.Context) {
	var service service.CreateArticleService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteArticle 删除微博内容
func DeleteArticle(c *gin.Context) {
	var service service.DeleteArticleService
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}


// ListArticle 查询单个微博内容
func ListArticle(c *gin.Context) {
	service := service.ListArticleService{}
	res := service.List(c.Param("id"))
	c.JSON(200, res)
}


func ShowArticles(c *gin.Context) {
	service := service.ShowArticlesService{}
	if err := c.ShouldBind(&service); err == nil{
		res := service.Show()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UpdateArticle(c *gin.Context) {
	service := service.UpdateArticleService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GenerateArticlePoster (c *gin.Context) {
	article := &service.Article{}
	qr := qrcode.NewQrcode(QRCODE_URL, 300, 300, qr.M, qr.Auto)
	posterName := service.GetPosterFlag() + "-" + qrcode.GetQrCodeFileName(qr.URL) + qr.GetQrCodeExt()
	articlePoster := service.NewArticlePoster(posterName, article, qr)
	articlePosterBgService := service.NewArticlePosterBg(
		"bg.jpg",
		articlePoster,
		&service.Rect{
			X0: 0,
			Y0: 0,
			X1: 550,
			Y1: 700,
		},
		&service.Pt{
			X: 125,
			Y: 298,
		},
	)

	_, filePath, err := articlePosterBgService.Generate()
	if err != nil {
		c.JSON(200, ErrorResponse(err))
	}

	c.JSON(200, serializer.Response{
		Msg: "上传成功",
		Data: map[string]string{
			"poster_url":      qrcode.GetQrCodeFullUrl(posterName),
			"poster_save_url": filePath + posterName,
		},
	})
}