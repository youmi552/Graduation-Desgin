package controller

import (
	"GraduationDesign/e"
	"GraduationDesign/models"
	"GraduationDesign/service"
	"GraduationDesign/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userService service.UserService

// 用户注册
func Register(c *gin.Context) {
	//获取前端的数据
	var rInfo models.RegisterInfo
	err := c.ShouldBind(&rInfo)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}

	//检验前端的数据
	//检验用户名是否为空
	if rInfo.UserName == "" {
		c.JSON(http.StatusOK, Resp(e.USERNAMEISNULL, nil))
		return
	}
	//检验用户名长度
	if len(rInfo.UserName) > 19 {
		c.JSON(http.StatusOK, Resp(e.INVAILDUSERNAME, nil))
		return
	}
	//检验密码是否为空
	if rInfo.Password == "" {
		c.JSON(http.StatusOK, Resp(e.PASSWORDISNULL, nil))
		return
	}
	//检验密码长度
	if len(rInfo.Password) < 6 || len(rInfo.Password) > 12 {
		c.JSON(http.StatusOK, Resp(e.INVAILDPASSWORD, nil))
		return
	}
	//检验验证码是否为空
	if rInfo.Code == "" {
		c.JSON(http.StatusOK, Resp(e.CODEISNULL, nil))
		return
	}
	//检验手机号码的真实性
	err = util.IsMobile(rInfo.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.INVAILDPHONENUMBER, nil))
		return
	}
	//调用服务层注册的接口
	err = userService.Register(rInfo)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.REGISTERFAIL, err.Error()))
		return
	}
	//以json的数据格式传入给前端
	c.JSON(http.StatusOK, Resp(e.SUCCESS, nil))
}

// 邮箱注册
//func Register2(c *gin.Context) {
//	//获取前端的数据
//	var user models.User
//	code, _ := c.GetPostForm("code")
//	err := c.ShouldBind(&user)
//	if err != nil {
//		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
//		return
//	}
//
//	//检验前端的数据
//	//检验用户名是否为空
//	if user.UserName == "" {
//		c.JSON(http.StatusOK, Resp(e.USERNAMEISNULL, nil))
//		return
//	}
//	//检验密码是否为空
//	if user.Password == "" {
//		c.JSON(http.StatusOK, Resp(e.PASSWORDISNULL, nil))
//		return
//	}
//	//检验密码长度
//	if len(user.Password) < 6 || len(user.Password) > 12 {
//		c.JSON(http.StatusOK, Resp(e.INVAILDPASSWORD, nil))
//		return
//	}
//	//检验验证码是否为空
//	if code == "" {
//		c.JSON(http.StatusOK, Resp(e.CODEISNULL, nil))
//		return
//	}
//	//调用服务层注册的接口
//	err = userService.Register2(user, code)
//	if err != nil {
//		c.JSON(http.StatusOK, Resp(e.REGISTERFAIL, err.Error()))
//		return
//	}
//	//以json的数据格式传入给前端
//	c.JSON(http.StatusOK, Resp(e.SUCCESS, nil))
//}

// 获取验证码
func GetVerification(c *gin.Context) {
	//获取前端传来的参数
	phoneNumber := c.Param("phonenumber")
	fmt.Println(len(phoneNumber))
	//检验手机号码的真实性
	err := util.IsMobile(phoneNumber)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.INVAILDPHONENUMBER, err.Error()))
		return
	}
	//转交给服务层获取验证码
	err = userService.GetVerificationCode(phoneNumber)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETVERIFICATIONCODEFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, nil))
}

// 获取验证码
//func GetVerification2(c *gin.Context) {
//	//获取前端传来的参数
//	email := c.Param("email")
//	//转交给服务层获取验证码
//	err := userService.GetVerificationCode2(email)
//	if err != nil {
//		c.JSON(http.StatusOK, Resp(e.GETVERIFICATIONCODEFAIL, err.Error()))
//		return
//	}
//	c.JSON(http.StatusOK, Resp(e.SUCCESS, nil))
//}

// 用户登录
func Login(c *gin.Context) {
	//获取前端传来的参数
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}
	//检验用户名是否为空
	if user.UserName == "" {
		c.JSON(http.StatusOK, Resp(e.USERNAMEISNULL, nil))
		return
	}
	//检验用户密码是否为空
	if user.Password == "" {
		c.JSON(http.StatusOK, Resp(e.PASSWORDISNULL, nil))
		return
	}
	//转交给业务层登录
	lresp, err := userService.Login(user)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.LOGINFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, lresp))

}

// 用户电话登录
func LoginByPhoneNumber(c *gin.Context) {
	//获取前端传来的参数
	var loginInfo models.LoginInfo
	err := c.ShouldBind(&loginInfo)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.GETPARAMFAIL, nil))
		return
	}
	fmt.Println(loginInfo)
	//检验手机号码的真实性
	err = util.IsMobile(loginInfo.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.INVAILDPHONENUMBER, nil))
		return
	}
	//转交给服务层进行登录
	lresp, err := userService.LoginByPhoneNumber(loginInfo.PhoneNumber, loginInfo.Code)
	if err != nil {
		c.JSON(http.StatusOK, Resp(e.LOGINFAIL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, Resp(e.SUCCESS, lresp))
}
