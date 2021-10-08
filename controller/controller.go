package controller

import (
	"gin_demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context)  {
	// 从请求中把数据取出来
	var todo models.Todo
	c.BindJSON(&todo)
	// 存入数据库
	err := models.CreateTodo(&todo)
	// 相应
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context)  {
	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context)  {
	id, ok := c.Params.Get("id")
	if ! ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "err",
		})
	} else {
		todo, err := models.GetTodoById(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			c.BindJSON(&todo)
			err = models.UpdateTodo(todo)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		}
	}
}

func DeleteTodo(c *gin.Context)  {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "err",
		})
	} else {
		var todo models.Todo
		err := models.DeleteTodo(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, todo)
		}
	}
}