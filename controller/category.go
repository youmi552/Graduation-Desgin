package controller

import (
	"GraduationDesign/e"
	"GraduationDesign/models"
	"GraduationDesign/service"
	"github.com/gin-gonic/gin"
)

var categoryService service.CategoryService

func AddCategory(c *gin.Context) {
	var category models.CategoryParam
	err := c.ShouldBind(&category)
	if err != nil {
		c.JSON(200, Resp(e.INVAILDPARAMS, err.Error()))
		return
	}
	err = categoryService.AddCategory(category)
	if err != nil {
		c.JSON(200, Resp(e.ADDCATEGORYFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, nil))

}
func DeleteCategory(c *gin.Context) {
	var category models.CategoryParam
	err := c.ShouldBind(&category)
	if err != nil {
		c.JSON(200, Resp(e.INVAILDPARAMS, err.Error()))
		return
	}
	err = categoryService.DeleteCategory(category)
	if err != nil {
		c.JSON(200, Resp(e.DELETECATEGORYFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, nil))
}

func GetAllCategory(c *gin.Context) {
	categorys, err := categoryService.GetAllCategory()
	if err != nil {
		c.JSON(200, Resp(e.GETALLCATEGORYFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, categorys))
}

func GetAllAdviceCategory(c *gin.Context) {
	categorys, err := categoryService.GetAllAdviceCategory()
	if err != nil {
		c.JSON(200, Resp(e.GETALLCATEGORYFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, categorys))
}
