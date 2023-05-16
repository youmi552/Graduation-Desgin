package middleware

import (
	"GraduationDesign/controller"
	"GraduationDesign/dao"
	"GraduationDesign/e"
	"GraduationDesign/models"
	"GraduationDesign/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userDao dao.UserDao

// 检验用户token有效期
func LoginNeed(c *gin.Context) {
	//获取token
	token := c.Request.Header.Get("Authorization")
	//解密token获得用户uid
	_, claim, err := util.ParseToken(token)
	//解密失败则返回登录已过期提示
	if err != nil || claim.Uid < 1 {
		c.JSON(http.StatusOK, controller.Resp(e.NEEDLOGIN, err.Error()))
		c.Abort()
		return
	}
	//将token交给服务层检验用户token是否为最新
	token2, err := userDao.GetTokenByUid(claim.Uid)
	if err != nil {
		c.JSON(http.StatusOK, &models.Resp{
			Code: e.NEEDLOGIN,
			Msg:  e.GetMsg(e.NEEDLOGIN),
			Data: nil,
		})
		c.Abort()
		return
	}
	//检验登录设备是否唯一
	if token != token2 {
		c.JSON(http.StatusOK, &models.Resp{
			Code: e.TOKENERROR,
			Msg:  e.GetMsg(e.TOKENERROR),
			Data: nil,
		})
		c.Abort()
		return
	}
	c.Set("uid", claim.Uid)
	//放行
	c.Next()
}
