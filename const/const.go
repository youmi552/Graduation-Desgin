package _const

import "github.com/spf13/viper"

var SeverAddress string
var AvatarMin int
var AvatarMax int

func InitConst() {
	SeverAddress = viper.GetString("server.address")
	AvatarMin = viper.GetInt("avatar.min")
	AvatarMax = viper.GetInt("avatar.max")
}
