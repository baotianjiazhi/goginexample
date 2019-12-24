package api

import (
	"ginexample/service"
	"github.com/gin-gonic/gin"
)

func CreateTag(c *gin.Context) {
	var service service.CreateTagService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateTag()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


func DeleteTag(c *gin.Context) {
	var service service.DeleteTagService
	res := service.DeleteTag(c.Param("id"))
	c.JSON(200, res)
}

func UpdateTag(c *gin.Context) {
	var service service.UpdateTagService
	if  err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}