package models

type Resp struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type LoginResp struct {
	UserInfo UserInfo `json:"userInfo"`
	Token    string   `json:"token"`
}
