package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
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

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := initMySQL()
	if err != nil {
		panic("failed to connect database")
	}
	// 程序退出关闭数据库连接
	defer Close()

	// 将模型与数据库中的表绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	// 加载静态模板，模板文件引用的静态文件位置
	r.Static("/static", "./dist/static")
	// 模板解析,去哪里找模板文件
	// r.LoadHTMLFiles("dist/index.html", "dist/favicon.ico")
	r.LoadHTMLGlob("dist/template/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 前端页面填写待办事项，点击提交，发送请求至此
			// 从请求中拿出数据
			var todo Todo
			c.BindJSON(&todo)
			// 将数据存入数据库
			// 返回响应
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 2000,
					"msg":  "success",
					"data": todo,
				})
			}
		})
		// 查看所有事项
		v1Group.GET("/todo", func(c *gin.Context) {
			// 查询todos表中的所有数据
			var todolist []Todo
			if err = DB.Find(&todolist).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todolist)
			}
		})
		// 查看某一事项，利用url参数
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "id无效"})
				return
			}
			var todo Todo
			if err = DB.Where("id = ?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 修改某一事项，利用url参数
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			// 获取需要更新的事项的id
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "id无效"})
				return
			}
			var todo Todo
			// 在数据库中查找指定id的数据
			if err = DB.Where("id = ?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			// 接收从前端传来的更新数据值
			c.BindJSON(&todo)
			// 更新数据库中该条数据的值
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 删除某一事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			// 获取需要删除的事项的id
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "id无效"})
				return
			}
			// 在数据库中删除指定id的数据db.Delete(&Todo{}, 10)
			if err = DB.Delete(&Todo{}, id).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"id": "deleted",
				})
			}
		})
	}

	r.Run(":9090")
}
