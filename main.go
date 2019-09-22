package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ranggakusuma/go-todo/controllers"
)

// Router is global variable for gin engine
var Router *gin.Engine

func main() {

	Routes()
	Router.Run()
	fmt.Println("Running server")
}

// Routes for run routes
func Routes() {
	Router = gin.Default()
	var todoController controllers.TodoStruct
	todoController.Routes(Router)
}
