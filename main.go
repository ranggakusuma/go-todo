package main

import (
	"fmt"

	"github.com/ranggakusuma/go-todo/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("test")
	var todoController controllers.TodoStruct
	router := gin.Default()

	v1 := router.Group("api/todos")

	v1.GET("/", todoController.Get)
	v1.GET("/:id", todoController.One)
	v1.POST("/", todoController.Create)

	router.Run()

	fmt.Println("Running server")
}
