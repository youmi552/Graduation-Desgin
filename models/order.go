package models

import "time"

type Order struct {
	Oid    int    `json:"oid" gorm:"primarykey" form:"oid"` //订单id
	Status int    `json:"status"`                           //订单状态
	Uid    int    `json:"uid"`                              //买家uid
	Sid    int    `json:"sid" form:"sid"`                   //卖家id
	Gid    int    `json:"gid"  form:"gid"`                  //物品id
	Place  string `json:"place" form:"place"`               //交易地址
	//Price     float32   `json:"price"`                 //订单价格
	Notes     string    `json:"notes"  form:"notes"` //备注
	CreatedAt time.Time `json:"createdAt"`           //创建时间
	UpdatedAt time.Time `json:"updatedAt"`           //更新时间
}
type OrderGid struct {
	Oid int //订单id
	Gid int //物品id
}
type OrderInfo struct {
	Oid          int       `json:"oid" gorm:"primarykey"` //订单id
	Goods        Goods     `json:"goodsData"`             //物品信息
	UserName     string    `json:"userName"`              //买家用户名
	SUserName    string    `json:"sUserName"`
	PhoneNumber  string    `json:"phoneNumber"` //买家联系电话
	SPhoneNumber string    `json:"sPhoneNumber"`
	Status       int       `json:"status"`    //订单状态
	Place        string    `json:"place"`     //交易地址
	Price        float32   `json:"price"`     //订单价格
	Notes        string    `json:"notes"`     //备注
	CreatedAt    time.Time `json:"createdAt"` //创建时间
	UpdatedAt    time.Time `json:"updatedAt"` //更新时间
	GoodsInfo    GoodsInfo `json:"goodsInfo"`
}
