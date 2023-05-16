package middleware

import (
	controller "GraduationDesign/controller"
	"GraduationDesign/e"
	"GraduationDesign/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminNeed(c *gin.Context) {
	uid, _ := c.Get("uid")
	uid2 := util.GetInterfaceToString(uid)
	//查询用户权限等级
	level, err := userDao.GetUserLevelByUid(uid2)
	if err != nil {
		c.JSON(http.StatusOK, controller.Resp(e.INVAILDPARAMS, err.Error()))
		c.Abort()
	}
	if level != 2 {
		c.JSON(http.StatusOK, controller.Resp(e.ADMINNEED, nil))
		c.Abort()
	}
	c.Next()
}
