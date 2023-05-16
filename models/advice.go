package models

import "time"

type Advice struct {
	Aid         int `json:"aid" gorm:"primarykey"`
	Uid         int
	Cid         int       `json:"cid" form:"categoryId"` //分类id
	Title       string    `json:"title" form:"title"`
	PhoneNumber string    `json:"phoneNumber" form:"phoneNumber"`
	Detail      string    `json:"detail" form:"detail"`
	Status      int       `json:"status" form:"status"`
	CreatedAt   time.Time `json:"createdAt"` //创建时间
	UpdatedAt   time.Time `json:"updatedAt"` //更新时间
}
type AdviceData struct {
	Aid       int    `json:"aid" gorm:"primarykey"`
	UserName  string `json:"userName"`
	Category  string `json:"category"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	CreatedAt string `json:"createdAt"` //创建时间
}
type AdviceDetail struct {
	Aid         int      `json:"aid" gorm:"primarykey"`
	UserName    string   `json:"userName"`
	Category    string   `json:"category"`
	Title       string   `json:"title"`
	Status      int      `json:"status"`
	Pictures    []string `json:"pictures"`
	PhoneNumber string   `json:"phoneNumber" form:"phoneNumber"`
	Detail      string   `json:"detail" form:"detail"`
	CreatedAt   string   `json:"createdAt"` //创建时间
}
type Photo struct {
	Pid   int    `gorm:"primarykey"` //图片id
	Aid   int    //商品对应的图片
	Photo string //图片地址
}
