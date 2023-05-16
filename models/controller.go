package models

import "mime/multipart"

type PageParams struct {
	PageNumber int    `form:"pagenumber"`
	PageSize   int    `form:"pagesize"`
	Keyword    string `form:"keyword"`
}

// 登录信息
type LoginInfo struct {
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Code        string `json:"code" form:"code"`
}

// 注册信息
type RegisterInfo struct {
	UserName    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Code        string `json:"code" form:"code"`
}

// 文件信息
type FileParams struct {
	Gid string                `json:"gid" form:"gid"`
	F1  *multipart.FileHeader `json:"f1" form:"f1"`
}

// 收藏参数
type CollectionParams struct {
	Gid          string `json:"gid" form:"gid"`
	IsCollection bool   `json:"isCollection" form:"isCollection"`
}

type CategoryParams struct {
	Cid        int `form:"cid"`
	PageNumber int `form:"pagenumber"`
	PageSize   int `form:"pagesize"`
}
