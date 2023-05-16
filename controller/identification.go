package controller

import (
	"GraduationDesign/e"
	"GraduationDesign/models"
	"GraduationDesign/service"
	"GraduationDesign/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var identificationService service.IdentificationService

func GetIdentification(c *gin.Context) {
	//获取用户的登录用户的uid
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, nil))
		return
	}
	status, err := identificationService.GetIdentificationByUid(uid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETIDENTIFICATIONFAIL, nil))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, status))
}

func AddIdentification(c *gin.Context) {
	//获取用户的登录用户的uid
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, nil))
		return
	}
	var identification models.Identification
	err = c.ShouldBind(&identification)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}
	identification.Uid = uid
	err = identificationService.AddIdentification(identification)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.ADDIDENTIFICATIONFAIL, nil))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, nil))

}
func UpdateIdentification(c *gin.Context) {
	//获取用户的登录用户的uid
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, nil))
		return
	}
	var identification models.Identification
	err = c.ShouldBind(&identification)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}
	identification.Uid = uid
	err = identificationService.UpdateIdentification(identification)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.UPDATEIDENTIFICATIONFAIL, nil))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, nil))
}

func GetAllIdentification(c *gin.Context) {
	var pageParams models.PageParams
	err := c.ShouldBind(&pageParams)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}
	identifications, err := identificationService.GetAllIdentification(pageParams)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETIDENTIFICATIONFAIL, nil))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, identifications))
}

func AcceptIdentification(c *gin.Context) {
	i := c.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, nil))
		return
	}
	err = identificationService.AcceptIdentification(id)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.ACCEPTIDENTIFICATIONFAIL, nil))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, nil))
}
func RefuseIdentification(c *gin.Context) {
	i := c.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, nil))
		return
	}
	err = identificationService.RefuseIdentification(id)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.REFUSEIDENTIFICATIONFAIL, nil))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, nil))
}
