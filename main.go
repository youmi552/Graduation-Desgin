package main

import (
	"GraduationDesign/config"
	_const "GraduationDesign/const"
	"GraduationDesign/dao"
	"GraduationDesign/router"
	"fmt"
)

func init() {
	//初始化配置文件
	config.InitConfig()
	//初始化常量
	_const.InitConst()
	//初始化mysql
	dao.InitMysql()
	//初始化redis
	dao.InitRedis()
	//初始化路由
	router.InitRouter()

}

func main() {
	fmt.Println("项目已启动!")
}
