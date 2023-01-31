package main

import (
	"ginlist/database"
	"ginlist/model"
	"ginlist/router"
)

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := database.InitMySQL()
	if err != nil {
		panic("failed to connect database")
	}
	// 程序退出关闭数据库连接
	defer database.Close()

	// 将模型与数据库中的表绑定
	database.DB.AutoMigrate(&model.Todo{})

	r := router.SetRouter()
	r.Run(":9090")
}
