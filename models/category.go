package models

type Category struct {
	Cid   string `json:"cid"  form:"cid"`
	CName string `json:"CName"  form:"cName"`
}
type CategoryParam struct {
	Cid   string `json:"cid"  form:"cid"`
	CName string `json:"CName"  form:"cName"`
	Type  string `json:"type" form:"type"`
}
