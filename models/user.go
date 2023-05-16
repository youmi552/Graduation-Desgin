package models

import "time"

// 用户
type User struct {
	Uid      int    `gorm:"primarykey" column:"uid" form:"uid"`     //用户id
	UserName string `form:"username" json:"username" gorm:"unique"` //用户名
	Password string `form:"password" json:"password"`               //密码
	//Email        string    `form:"email"`                                                 //邮箱
	Avatar       string    `json:"avatar"`                                       //用户头像
	PhoneNumber  string    `form:"phonenumber" json:"phoneNumber" gorm:"unique"` //电话号码
	Introduction string    `json:"introduction" gorm:"default:这个人很懒,什么都没有留下"`    //个人简介
	Token        string    `json:"token"`                                        //登录凭证
	Level        int       `json:"level" gorm:"default:0"`                       //用户级别
	CreatedAt    time.Time `json:"createdAt"`                                    //创建时间
	UpdatedAt    time.Time `json:"updatedAt"`                                    //更新时间
}

// 用户信息
type UserInfo struct {
	Uid          int    `json:"uid" gorm:"uid"`                   //用户id
	UserName     string `json:"username" gorm:"user_name"`        //用户名
	Avatar       string `json:"avatar" gorm:"avatar"`             //用户头像
	Introduction string `json:"introduction" gorm:"introduction"` //个人简介
	Level        int    `json:"level" goem:"level"`               //用户级别
}
