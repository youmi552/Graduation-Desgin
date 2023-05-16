package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitConfig() {
	//获得工作路径
	workDir, _ := os.Getwd()
	//设置配置文件名称
	viper.SetConfigName("config")
	//设置配置文件类型
	viper.SetConfigType("yml")
	//添加配置文件加载路径
	viper.AddConfigPath(workDir + "/config")
	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("加载配置文件出错！")
		return
	}
	log.Println("配置文件加载成功。")
}
func InitConfigTest() {
	//获得工作路径
	workDir, _ := os.Getwd()
	//设置配置文件名称
	viper.SetConfigName("config")
	//设置配置文件类型
	viper.SetConfigType("yml")
	//添加配置文件加载路径
	viper.AddConfigPath(workDir + "/")
	fmt.Println(workDir)
	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("加载配置文件出错！")
		return
	}
	log.Println("配置文件加载成功。")
}
