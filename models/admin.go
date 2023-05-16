package models

type UserData struct {
	Uid         int    `json:"uid"`
	UserName    string `json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
	Level       int    `json:"level"`
}
