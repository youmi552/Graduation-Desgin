package dao

import (
	"GraduationDesign/models"
	"fmt"
)

type UserInfoDao struct {
}

// 修改用户头像
func (*UserInfoDao) UpdateAvatar(uid int, avatar string) error {
	user := models.User{Avatar: avatar}
	tx := db.Model(&models.User{Uid: uid}).Updates(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// 修改用户个人简介
func (*UserInfoDao) UpdateIntroduction(uid int, introduction string) error {
	user := models.User{Introduction: introduction}
	tx := db.Model(&models.User{Uid: uid}).Updates(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// 获取用户信息
func (*UserInfoDao) GetUserByUsername(username string) (int, error) {
	var uid int
	tx := db.Model(&models.User{}).Select("uid").Where("user_name=?", username).First(&uid)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return uid, nil
}

// 修改用户名
func (*UserInfoDao) UpdateUsername(uid int, username string) error {
	user := models.User{UserName: username}
	tx := db.Model(&models.User{Uid: uid}).Updates(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// 修改用户信息
func (*UserInfoDao) GetUserInfo(uid int) (models.UserInfo, error) {
	var userInfo models.UserInfo
	fmt.Println(uid)
	tx := db.Model(&models.User{}).Where("uid=?", uid).First(&userInfo)
	if tx.Error != nil {
		return userInfo, tx.Error
	}
	return userInfo, nil

}

// 通过uid获取用户信息
func (*UserInfoDao) GetUserNameByUid(uid int) (string, error) {
	var userName string
	tx := db.Model(&models.User{}).Select("user_name").Where("uid=?", uid).First(&userName)
	if tx.Error != nil {
		return "", tx.Error
	}
	return userName, nil
}

// 获取所有用户信息
func (*UserInfoDao) GetAllUserInfo() ([]models.UserInfo, error) {
	var userInfos []models.UserInfo
	tx := db.Model(&models.User{}).Where("level=0").Find(&userInfos)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return userInfos, nil
}
