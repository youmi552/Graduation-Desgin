package service

import (
	_const "GraduationDesign/const"
	"GraduationDesign/dao"
	"GraduationDesign/models"
	"errors"
	"fmt"
	"strconv"
)

type UserInfoService struct {
}

var userInfoDao dao.UserInfoDao

// 更新用户头像
func (*UserInfoService) UpdateAvatar(uid int, numberAvatar string) (models.UserInfo, error) {
	avatarNumber, _ := strconv.Atoi(numberAvatar)
	if avatarNumber < _const.AvatarMin || avatarNumber > _const.AvatarMax {
		return models.UserInfo{}, errors.New("头像参数错误！")
	}
	avatar := fmt.Sprint("http://"+_const.SeverAddress, "/file/系统头像", numberAvatar, ".jpg")
	err := userInfoDao.UpdateAvatar(uid, avatar)
	userInfo, err := userInfoDao.GetUserInfo(uid)
	if err != nil {
		return models.UserInfo{}, err
	}
	return userInfo, nil
}

// 更新用户简介
func (*UserInfoService) UpdateIntroduction(uid int, introduction string) error {
	err := userInfoDao.UpdateIntroduction(uid, introduction)
	if err != nil {
		return err
	}
	return nil
}

// 更新用户名
func (*UserInfoService) UpdateUsername(uid int, username string) error {
	userid, _ := userInfoDao.GetUserByUsername(username)
	if userid >= 1 {
		if userid == uid {
			return errors.New("已经是该用户名！")
		}
		return errors.New("用户名已经被使用！")
	}
	err := userInfoDao.UpdateUsername(uid, username)
	if err != nil {
		return err
	}
	return err
}

func (*UserInfoService) GetUserInfo(uid int) (models.UserInfo, error) {
	userInfo, err := userInfoDao.GetUserInfo(uid)
	if err != nil {
		return userInfo, err
	}
	return userInfo, nil
}

func (*UserInfoService) GetAllUserInfo() ([]models.UserInfo, error) {
	userInfos, err := userInfoDao.GetAllUserInfo()
	if err != nil {
		return nil, err
	}
	return userInfos, nil
}

func (*UserInfoService) UpdateUserData(user models.User) error {
	err := userDao.UpdateUserData(user)
	if err != nil {
		return err
	}
	return nil
}
