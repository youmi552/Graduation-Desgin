package service

import (
	_const "GraduationDesign/const"
	"GraduationDesign/dao"
	"GraduationDesign/models"
	"GraduationDesign/util"
	"fmt"
	"log"
)

type UserService struct {
}

var userDao dao.UserDao

// 用户注册
func (userService *UserService) Register(rInfo models.RegisterInfo) error {
	var user = models.User{
		UserName:    rInfo.UserName,
		Password:    rInfo.Password,
		PhoneNumber: rInfo.PhoneNumber,
	}
	//验证验证码
	bool, err := userService.verification(user.PhoneNumber, rInfo.Code)
	if err != nil {
		return err
	}
	if bool == false {
		return err
	}
	//加密密码
	user.Password = util.Md5Crypt(user.Password, "youmigd")
	avatar := fmt.Sprint(_const.SeverAddress, "/file/系统头像0.jpg")
	user.Avatar = avatar
	//进行注册
	err = userDao.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

//// 邮箱注册
//func (userService *UserService) Register2(user models.User, code string) error {
//	//验证验证码
//	bool, err := userService.verification(user.Email, code)
//	if err != nil {
//		return err
//	}
//	if bool == false {
//		return err
//	}
//	//加密密码
//	user.Password = util.Md5Crypt(user.Password, "youmigd")
//	//进行注册
//	err = userDao.CreateUser(user)
//	if err != nil {
//		return err
//	}
//	return nil
//}

// 获取验证码,时效为5分钟
func (userService *UserService) GetVerificationCode(phoneNumber string) error {
	//创建验证码
	code := util.Random5()
	//将验证码存入
	err := userDao.CreateVerificationCode(phoneNumber, code)
	if err != nil {
		return err
	}
	//向用户发送验证码
	log.Println("验证码为:", code)
	return err
}

// 获取验证码,时效为5分钟
//func (userService *UserService) GetVerificationCode2(email string) error {
//	//创建验证码
//	code := util.Random5()
//	//向用户发送验证码
//	m := gomail.NewMessage()
//	//发送人
//	m.SetHeader("From", "2229262420@qq.com")
//	//接收人
//	m.SetHeader("To", email)
//	//抄送人
//	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
//	//主题
//	m.SetHeader("Subject", "小佩奇")
//	//内容
//	m.SetBody("text/html", "<h1>"+code+"</h1>")
//	//附件
//	//m.Attach("./myIpPic.png")
//	//拿到token，并进行连接,第4个参数是填授权码
//	d := gomail.NewDialer("smtp.qq.com", 587, "2229262420@qq.com", "akjjjoztopscebij")
//	// 发送邮件
//	if err := d.DialAndSend(m); err != nil {
//		fmt.Printf("DialAndSend err %v:", err)
//		panic(err)
//	}
//	fmt.Println("send mail success\n")
//	//向用户发送验证码
//	log.Println("验证码为:", code)
//	//将验证码存入
//	err := userDao.CreateVerificationCode(email, code)
//	if err != nil {
//		return err
//	}
//	return nil
//}

// 验证验证码是否正确
func (userService *UserService) verification(phoneNumber string, code string) (bool, error) {
	//获取redis中的验证码信息
	code2, err := userDao.GetVerificationCode(phoneNumber)
	if err != nil {
		return false, err
	}
	if code == code2 {
		return true, nil
	} else {
		return false, nil
	}
}

// 通过电话号码登录
func (userService *UserService) LoginByPhoneNumber(phoneNumber string, code string) (*models.LoginResp, error) {
	//查询数据库中是否有该用户
	userInfo, err := userDao.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		return &models.LoginResp{}, err
	}
	//验证验证码
	bool, err := userService.verification(phoneNumber, code)
	if err != nil {
		return &models.LoginResp{}, err
	}
	if bool == false {
		return &models.LoginResp{}, err
	}
	//生成token
	token, err := util.Award(&userInfo.Uid)
	if err != nil {
		return &models.LoginResp{}, err
	}
	err = userDao.SetToken(userInfo.Uid, token)
	if err != nil {
		return &models.LoginResp{}, err
	}
	//返回登录信息给控制层
	return &models.LoginResp{
		UserInfo: userInfo,
		Token:    token,
	}, nil
}

// 登录
func (userService *UserService) Login(user models.User) (*models.LoginResp, error) {
	//加密用户输入的密码
	user.Password = util.Md5Crypt(user.Password, "youmigd")
	//查询数据库中是否有该用户
	user2, err := userDao.GetUserByUserNameAndPassword(user)
	if err != nil {
		return nil, err
	}
	//检验用户名输入的密码
	if user.Password != user2.Password {
		return nil, err
	}
	//生成token
	token, err := util.Award(&user2.Uid)
	if err != nil {
		return nil, err
	}
	err = userDao.SetToken(user2.Uid, token)
	if err != nil {
		return nil, err
	}
	//返回登录信息给控制层
	return &models.LoginResp{
		UserInfo: models.UserInfo{Uid: user2.Uid, UserName: user2.UserName, Avatar: user2.Avatar, Introduction: user2.Introduction, Level: user2.Level},
		Token:    token,
	}, nil

}
