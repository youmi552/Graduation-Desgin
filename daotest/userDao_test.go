package daotest

import (
	"GraduationDesign/config"
	"GraduationDesign/dao"
	"GraduationDesign/service"
	"fmt"
	"os"
	"testing"
	"time"
)

func init() {
	config.InitConfigTest()
	dao.InitMysql()
	dao.InitRedis()
}

var userDao dao.UserDao
var userService service.UserService

func TestWorkDir(t *testing.T) {
	dir, _ := os.Getwd()
	fmt.Println(dir)
}

func TestRegister(t *testing.T) {
	//fmt.Println("——————————————————————————————")
	//fmt.Println("测试UserDao注册接口")
	//resp, err := userService.Register(models.User{}, "6465")
	//if err != nil || resp.Code != 200 {
	//	log.Println("注册失败！2秒后退出测试。错误信息:", resp)
	//	time.Sleep(time.Second * 2)
	//	return
	//}
	//fmt.Println("注册成功！,返回给前端信息:", resp)
}
func TestGetVerificationCode(t *testing.T) {
	fmt.Println("——————————————————————————————")
	fmt.Println("测试UserService获取验证码接口")
	userService.GetVerificationCode("18782436928")
	return
}
func TestGetUserInfoByPhoneNumber(t *testing.T) {
	fmt.Println("——————————————————————————————")
	fmt.Println("测试UserDao通过电话号码获取用户信息接口")
	userInfo, err := userDao.GetUserByPhoneNumber("18782436928")
	if err != nil {
		fmt.Println("没有该用户！")
	}
	fmt.Println(userInfo)
	return
}
func TestGetNotice(t *testing.T) {
	var noticeDao dao.NoticeDao
	notice, _ := noticeDao.GetNotice()
	//stringMap, _ := redis.String(notice, errors.New("niu"))
	fmt.Println(notice)
}
func TestSimple(t *testing.T) {
	now := time.Now()
	s := now.Format("2006-01-03")
	fmt.Println(s)
}
