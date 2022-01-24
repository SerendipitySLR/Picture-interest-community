package respository

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"ptc/internal/model"
)

var db *gorm.DB
var err error

//初始化数据库上下文
func InitDbContext() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("datasource.username"),
		viper.GetString("datasource.password"),
		viper.GetString("datasource.host"),
		viper.GetString("datasource.port"),
		viper.GetString("datasource.database"),
	)

	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 关闭外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		//退出程序
		os.Exit(1)
	}

	// 迁移数据表,还没统一，所以自己加自己的。。。
	_ = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.UserDetails{},
		&model.UserRelatedData{},
		&model.UserRegister{},
		&model.Collection{},
		&model.Comment{},
		&model.Follow{},
		&model.Like{},
		&model.Post{},
		&model.Feeds{},
		&model.Forward{},
		&model.PostPhoto{})
}

func GetDB() *gorm.DB {
	return db
}
