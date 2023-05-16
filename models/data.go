package models

type Data struct {
	UserNumber     int
	GoodsNumber    int
	AdviceNumber   int
	OrderNumber    int
	GoodsCategory  []Category     `json:"productCategory"`
	AdviceCategory []Category     `json:"adviceCategory"`
	CategoryData   []CategoryData `json:"categoryData"`
}

type CategoryData struct {
	Cid     string  `json:"cid"  form:"cid"`
	CName   string  `json:"CName"  form:"cName"`
	Percent float64 `json:"percent"`
}
