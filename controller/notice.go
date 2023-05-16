package controller

import (
	"GraduationDesign/e"
	"GraduationDesign/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var noticeService service.NoticeService

func GetNotice(c *gin.Context) {
	notice, err := noticeService.GetNotice()
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GetNoticeFail, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, notice))
}

// 修改公告
func UpdateNotice(c *gin.Context) {
	date, bool := c.GetPostForm("date")
	content, bool := c.GetPostForm("notice")
	if bool == false {
		c.JSON(http.StatusOK, Resp(e.INVAILDPARAMS, nil))
		return
	}
	notice, err := noticeService.UpdateNotice(date, content)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.UPDATENOTICEFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, notice))
}
