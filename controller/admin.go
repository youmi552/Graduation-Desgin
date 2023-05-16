package controller

import (
	"GraduationDesign/e"
	"GraduationDesign/models"
	"GraduationDesign/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var adminService service.AdminService

func GetData(c *gin.Context) {
	data, err := adminService.GetData()
	if err != nil {
		c.JSON(200, Resp(e.GETDATAFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, data))
}

func GetUserData(c *gin.Context) {
	userData, err := adminService.GetUserData()
	if err != nil {
		c.JSON(200, Resp(e.GETUSERDATAFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, userData))
}

func UpdateUserData(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.INVAILDPARAMS, err.Error()))
		return
	}
	err = adminService.UpdateUserData(user)
	if err != nil {
		c.JSON(200, Resp(e.UPDATEUSERDATAFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, nil))
}

func DeleteUser(c *gin.Context) {
	id := c.Param("uid")
	uid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	err = adminService.DeleteUser(uid)
	if err != nil {
		c.JSON(200, Resp(e.UPDATEUSERDATAFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, nil))
}

// 获得所有分类建议
func GetAllAdvice(c *gin.Context) {
	advice, err := adminService.GetAdvice()
	if err != nil {
		c.JSON(200, Resp(e.GETADVICEFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, advice))
}
func GetAdviceByAid(c *gin.Context) {
	id := c.Param("aid")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	advice, err := adviceService.GetAdviceByAid(aid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETADVICEFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, advice))

}

func ConfirmAdvice(c *gin.Context) {
	id := c.Param("aid")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	err = adminService.ConfirmAdvice(aid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.CONFIRMADVICEFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, nil))
}

func DeleteAdvice(c *gin.Context) {
	err := adminService.DeleteAdvice()
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.DELETEADVICEFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, nil))
}

func AllowGoods(c *gin.Context) {
	id := c.Param("gid")
	gid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	err = goodsService.UpdateGoodsStatus(gid, 1)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.ALLOWFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, nil))
}

func BanGoods(c *gin.Context) {
	id := c.Param("gid")
	gid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	err = goodsService.UpdateGoodsStatus(gid, 2)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.ALLOWFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, nil))
}

func GetAllGoodsData(c *gin.Context) {
	var pageParams models.PageParams
	err := c.ShouldBind(&pageParams)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}

	goods, err := adminService.GetGoodsByPage(pageParams.PageNumber, pageParams.PageSize, pageParams.Keyword)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETGOODSDATAFAIL, err))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, goods))
}
