package api

import (
	"ginexample/service"
	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	var service service.UploadImageService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Upload(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
