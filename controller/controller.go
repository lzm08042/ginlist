// 定义每个路由所绑定的执行处理过程
package controller

import (
	"ginlist/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 模板处理
func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// 添加，前端页面填写待办事项，点击提交，发送请求至此
func PostATodo(c *gin.Context) {
	// 从请求中拿出数据
	var todo model.Todo
	c.BindJSON(&todo)
	// 将数据存入数据库
	// 返回响应
	if err := model.CreateAItem(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": todo,
		})
	}
}

// 查看所有事项
func GetTodos(c *gin.Context) {
	// 查询todos表中的所有数据
	var todolist []*model.Todo
	var err error
	if todolist, err = model.FindItems(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

// 查看某一事项，利用url参数
func GetATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id无效"})
		return
	}
	if todo, err := model.FindAItem(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, &todo)
	}
}

// 修改某一事项，利用url参数
func PutATodo(c *gin.Context) {
	// 获取需要更新的事项的id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id无效"})
		return
	}

	// 在数据库中查找指定id的数据
	todo, err := model.FindAItem(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// 接收从前端传来的更新数据值
	c.BindJSON(&todo)
	// 更新数据库中该条数据的值
	if err := model.UpdateAItem(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// 删除某一事项
func DeleteATodo(c *gin.Context) {
	// 获取需要删除的事项的id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id无效"})
		return
	}
	// 在数据库中删除指定id的数据db.Delete(&Todo{}, 10)
	if err := model.DeleteAItem(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id": "deleted",
		})
	}
}
