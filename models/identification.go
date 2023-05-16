package models

import "time"

type Identification struct {
	Id        int       `json:"id" gorm:"primarykey"`
	Uid       int       `json:"uid"`
	Name      string    `json:"name" form:"name"`
	StudentId string    `json:"studentId" form:"studentId"`
	Status    int       `json:"status"`    //认证状态 0:待认证 1：认证成功 2：认证失败
	CreatedAt time.Time `json:"createdAt"` //创建时间
	UpdatedAt time.Time `json:"updatedAt"` //更新时间
}

type IdentificationData struct {
	Id          int    `json:"id" gorm:"primarykey"`
	UserName    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name" form:"name"`
	StudentId   string `json:"studentId" form:"studentId"`
	Status      int    `json:"status"` //认证状态 0:待认证 1：认证成功 2：认证失败
}

type IdentificationResp struct {
	Count              int                  `json:"count"`
	IdentificationData []IdentificationData `json:"identificationData"`
}
