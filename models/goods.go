package models

import (
	"time"
)

type Goods struct {
	Gid          int       `json:"gid" gorm:"primarykey"`            //商品id
	GoodsName    string    `json:"goodsName" form:"goodsName"`       //商品名称
	Uid          int       `json:"uid"`                              //所属用户id
	Cid          int       `json:"cid" form:"categoryId"`            //分类id
	Introduction string    `json:"introduction" form:"introduction"` //商品简介
	UsedTime     string    `json:"usedTime" form:"usedTime"`         //使用时间
	Price        float32   `json:"price" form:"price"`               //价格
	Status       int       `json:"status"`                           //商品状态 0:上架 1：下架
	CreatedAt    time.Time `json:"createdAt"`                        //创建时间
	UpdatedAt    time.Time `json:"updatedAt"`                        //更新时间
}
type GoodsInfo struct {
	Gid          int    `json:"gid" gorm:"primarykey"` //商品id
	GoodsName    string `json:"goodsName"`             //商品名称
	UserName     string `json:"userName"`              //所属用户名
	CategoryName string `json:"categoryName"`          //分类名
	Pictures     []string
	Introduction string  `json:"introduction"`             //商品简介
	UsedTime     string  `json:"usedTime"`                 //使用时间
	Price        float32 `json:"price"`                    //价格
	Status       int     `json:"status" gorm:"default:0" ` //商品状态 0:上架 1：下架
	Views        int     `json:"views"`                    //浏览量
	Collection   int     `json:"collection"`               //收藏量
	IsCollection bool    `json:"isCollection"`             //当前用户是否收藏 0:没有收藏 1：收藏了
}

type Pictures struct {
	Pid     int    `gorm:"primarykey"` //图片id
	Gid     int    //商品对应的图片
	Picture string //图片地址
}

// 商品展示页面数据
type GoodsData struct {
	Gid       int     `json:"gid" gorm:"primarykey"`      //商品id
	GoodsName string  `json:"goodsName" form:"goodsName"` //商品名称
	Picture   string  `json:"picture"`                    //展示图片
	UserName  string  `json:"userName"`                   //所属用户名
	UsedTime  string  `json:"usedTime" form:"usedTime"`   //使用时间
	Price     float32 `json:"price" form:"price"`         //价格
	Status    int     `json:"status"`
}

// 分页商品展示页面数据
type GoodsPageData struct {
	Count     int //数据总量
	GoodsData []GoodsData
}

// 分页商品展示页面数据
type GoodsPageDataWithAdmin struct {
	Ban       int
	Count     int //数据总量
	GoodsData []GoodsData
}
