// 进行数据的CURD操作
package model

import (
	"ginlist/database"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 将数据存入数据库
func CreateAItem(todo *Todo) (err error) {
	err = database.DB.Create(&todo).Error
	return
}

// 查询todos表中的所有数据
func FindItems() (todolist []*Todo, err error) {
	if err = database.DB.Find(&todolist).Error; err != nil {
		return nil, err
	}
	return
}

// 查看指定id对应的数据
func FindAItem(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = database.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

// 修改某一条数据
func UpdateAItem(todo *Todo) (err error) {
	err = database.DB.Save(todo).Error
	return
}

// 删除指定id对应的数据
func DeleteAItem(id string) (err error) {
	err = database.DB.Delete(&Todo{}, id).Error
	// err = database.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
