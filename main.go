package main

import (
	"fmt"
	"gin_demo/dao"
	"gin_demo/models"
	"gin_demo/routers"
)

func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect mysql success")
	}
	// 模型关联
	dao.DB.AutoMigrate(&models.Todo{})
	// 注册路由
	r := routers.SetRouter()
	r.Run(":8090")
}