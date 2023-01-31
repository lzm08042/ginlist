// 用于数据库等连接初始化
package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// 连接数据库
func InitMySQL() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	// 测试数据库连通性
	_, err = DB.DB()
	return
}

// Close 关闭数据库的连接，在main函数中以defer调用
func Close() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}
