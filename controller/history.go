package controller

import (
	"GraduationDesign/e"
	"GraduationDesign/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetHistory(c *gin.Context) {
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, err.Error()))
		return
	}
	goods, err := historyService.GetUserHistory(uid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETHISTORYFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, goods))
}
