package dao

import (
	"GraduationDesign/models"
	"GraduationDesign/util"
	"errors"
	"log"
)

type UserDao struct {
}

// 新建用户
func (*UserDao) CreateUser(user models.User) error {
	tx := db.Create(&user)
	if tx.Error != nil {
		log.Println("UserDao——Register:注册失败！")
		return errors.New("用户名已经被使用或该电话已经被用于注册！")
	}
	return nil
}

// 查找用户
func (*UserDao) Login(arg string, password string) (models.UserInfo, error) {
	var user models.User
	tx := db.Where("username=? or phone").First(&user)
	userInfo := models.UserInfo{
		Uid:          user.Uid,
		UserName:     user.UserName,
		Avatar:       user.Avatar,
		Introduction: user.Introduction,
	}
	if tx.Error != nil {
		log.Println("查找用户失败！")
		return userInfo, errors.New("查找用户失败！")
	}
	return userInfo, nil
}

// 生成验证码
func (*UserDao) CreateVerificationCode(phoneNumber string, code string) error {
	_, err := rd.Do("setex", phoneNumber, 300, code)
	if err != nil {
		log.Println("UserDao——CreateVerificationCode:创建验证码失败！")
		return errors.New("创建验证码失败！")
	}
	return nil
}

// 获取存储的验证码
func (*UserDao) GetVerificationCode(phoneNumber string) (string, error) {
	code, err := rd.Do("get", phoneNumber)
	if err != nil {
		return "", err
	}
	scode := util.GetInterfaceToString(code)
	if scode == "" {
		return "", errors.New("UserDao-GetVerification:请先获取验证码")
	}
	return scode, nil
}

// 通过电话获得用户信息
func (*UserDao) GetUserByPhoneNumber(phoneNumber string) (models.UserInfo, error) {
	var userInfo models.UserInfo
	tx := db.Model(&models.User{}).Where("phone_number=?", phoneNumber).First(&userInfo)
	if tx.Error != nil {
		return userInfo, tx.Error
	}
	return userInfo, nil
}

// 设置登录凭证
func (*UserDao) SetToken(uid int, token string) error {
	user := &models.User{Token: token}
	tx := db.Model(&models.User{}).Where("uid=?", uid).Updates(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// 获得用户信息
func (*UserDao) GetUserByUserNameAndPassword(user models.User) (models.User, error) {
	var user2 models.User
	tx := db.Model(&models.User{}).Where("user_name=? and password=?", user.UserName, user.Password).First(&user2)
	if tx.Error != nil {
		return user2, errors.New("没有该用户！")
	}
	return user2, nil
}

// 根据用户id获取token
func (*UserDao) GetTokenByUid(uid int) (string, error) {
	var token string
	tx := db.Model(&models.User{}).Select("token").Where("uid=?", uid).First(&token)
	if tx.Error != nil {
		return "", tx.Error
	}
	return token, nil
}

// 查询用户权限等级
func (*UserDao) GetUserLevelByUid(uid string) (int, error) {
	var level int
	tx := db.Model(&models.User{}).Select("level").Where("uid=?", uid).First(&level)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return level, nil
}

func (*UserDao) CountAllUsers() (int, error) {
	var count int64
	tx := db.Model(&models.User{}).Count(&count)
	if tx.Error != nil {
		log.Println(tx.Error)
		return 0, tx.Error
	}
	return int(count), nil
}

func (*UserDao) GetAllUser() ([]models.User, error) {
	var user []models.User
	tx := db.Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (*UserDao) UpdateUserData(user models.User) error {
	tx := db.Model(&models.User{}).Where("uid=?", user.Uid).Updates(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (*UserDao) DeleteUser(uid int) error {
	tx := db.Model(&models.User{}).Where("uid=?", uid).Delete(uid)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (*UserDao) GetUserByUid(uid int) (models.User, error) {
	var user models.User
	tx := db.Select("user_name", "phone_number").Where("uid=?", uid).First(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}
