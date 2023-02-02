// 封装路由列表
package router

import (
	"ginlist/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	// 加载静态模板，模板文件引用的静态文件位置
	r.Static("/static", "dist/static")
	// 模板解析,去哪里找模板文件
	// r.LoadHTMLFiles("dist/index.html", "dist/favicon.ico")
	r.LoadHTMLGlob("dist/template/*")
	r.GET("/v1", controller.GetIndex)

	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.PostATodo)
		// 查看所有事项
		v1Group.GET("/todo", controller.GetTodos)
		// 查看某一事项，利用url参数
		v1Group.GET("/todo/:id", controller.GetATodo)
		// 修改某一事项，利用url参数
		v1Group.PUT("/todo/:id", controller.PutATodo)
		// 删除某一事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
