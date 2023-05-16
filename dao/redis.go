package dao

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"log"
)

var rd redis.Conn

func InitRedis() {
	//获取配置文件中redis的地址
	address := viper.GetString("redis.address")
	//连接redis
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		log.Println("redis连接失败！")
		panic(err)
	}
	rd = conn
	log.Println("redis连接成功!")
}
