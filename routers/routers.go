package routers

import (
	"gin_demo/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	// load resource
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")

	// 首页
	r.GET("/", controller.IndexHandler)
	// 代办事项
	v1Group := r.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
