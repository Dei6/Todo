package models

import "gin_demo/dao"

type Todo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	if err != nil {
		return err
	}
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	}
	return todoList, nil
}

func GetTodoById(id string) (todo *Todo, err error) {
	todo = new(Todo)
	err = dao.DB.Where("id=?", id).Find(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}