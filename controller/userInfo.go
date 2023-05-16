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

var userInfoService service.UserInfoService

// 更新头像信息
func UpdateAvatar(c *gin.Context) {
	//获取用户的登录用户的uid
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, nil))
		return
	}
	//获取图片参数
	avatarNumber, bool := c.GetPostForm("avatarNumber")
	if !bool {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}
	//转交给服务层更新头像
	userInfo, err := userInfoService.UpdateAvatar(uid, avatarNumber)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.UPDATEAVATARERROR, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, userInfo))

}
func UpdateUserInfo(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.INVAILDPARAMS, err.Error()))
		return
	}
	//获取用户的登录用户的uid
	id, _ := c.Get("uid")
	user.Uid, err = strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, nil))
		return
	}
	err = userInfoService.UpdateUserData(user)
	if err != nil {
		c.JSON(200, Resp(e.UPDATEUSERDATAFAIL, err.Error()))
		return
	}
	c.JSON(200, Resp(e.SUCCESS, nil))
}

// 更新用户简介
func UpdateIntroduction(c *gin.Context) {
	//获取用户的登录用户的uid
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, nil))
		return
	}
	//获取用户简介信息
	introduction, bool := c.GetPostForm("introduction")
	if !bool {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}
	//验证用户简介是否为空
	if introduction == "" {
		introduction = "这个人很懒，什么都没留下。"
	}
	//转交给服务层更新头像
	err = userInfoService.UpdateIntroduction(uid, introduction)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.UPDATEINTRODUCTIONERROR, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, introduction))
}

// 更新用户名
func UpdateUsername(c *gin.Context) {
	//获取用户的登录用户的uid
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, nil))
		return
	}
	//获取用户名信息
	username, bool := c.GetPostForm("username")
	if !bool {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}
	//验证用户名是否为空
	if username == "" {
		c.JSON(http.StatusOK, Resp(e.USERNAMEISNULL, nil))
		return
	}
	//转交给服务层更新用户名
	err = userInfoService.UpdateUsername(uid, username)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.UPDATEUSERNAMEFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, username))
}
func GetUserInfo(c *gin.Context) {
	//获取用户的登录用户的uid
	id, _ := c.Get("uid")
	uid, err := strconv.Atoi(util.GetInterfaceToString(id))
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.STRCONVERROR, nil))
		return
	}
	//转交给业务层
	userInfo, err := userInfoService.GetUserInfo(uid)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETUSERINFOFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, userInfo))
}
func GetAllUserInfo(c *gin.Context) {
	userInfos, err := userInfoService.GetAllUserInfo()
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETUSERINFOFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(http.StatusOK, Resp(e.SUCCESS, userInfos)))
}
