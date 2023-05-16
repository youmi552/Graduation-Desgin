package dao

import (
	"GraduationDesign/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"strings"
)

var db *gorm.DB

func InitMysql() {
	//获取配置文件中的信息
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	charset := viper.GetString("mysql.charset")
	//设置mysql初始化DSN
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=", charset, "&parseTime=True&loc=Local"}, "")
	con, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 171,
	}), &gorm.Config{
		//是否跳过事务
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			//创建表添加前缀,User 会变为 gd_users
			TablePrefix: "gd_",
			//将表以单数形式存在, gd_users 会变为gd_user
			SingularTable: true,
		},
		//逻辑外键(代码里自动体现外键关系)
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Println("数据库加载失败！")
		panic(err)
	}
	db = con
	//DropTables()
	db.AutoMigrate(&models.Identification{})
	//db.AutoMigrate(&models.Goods{}, &models.Pictures{})
	//migrator := db.Migrator()
	//if migrator.HasTable(models.Category{}) == false {
	//	migrator.CreateTable(models.Category{})
	//}
	log.Println("数据库连接成功。")
}
func DropTables() {
	M := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").Migrator()
	M.DropTable(&models.User{})
}
