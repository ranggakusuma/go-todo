package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranggakusuma/go-todo/models"
	"github.com/ranggakusuma/go-todo/utils"
)

// TodoStruct struct
type TodoStruct struct{}

// func (t *TodoStruct) createTodo(c *gin.Context) {
// 	completed, _ := strconv.Atoi(c.PostForm("completed"))
// 	// todo := models.todoModel{Title: c.PostForm("title"), Completed: completed}
// 	todo := models.TodoModel{Title: c.}
// 	db.Save(&todo)
// 	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
// }

// // Todo is controller function for todos
// func Todo(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo item created successfully!"})
// }

// Create todo controller
func (t *TodoStruct) Create(c *gin.Context) {
	var todo models.Todo
	db := utils.DB()
	tx := db.Begin()
	// tx := db
	c.BindJSON(&todo)
	fmt.Println(todo)
	err := todo.Insert(tx)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error create Todo!"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{"message": "Todo item created successfully!"})
}

// One todo controller
func (t *TodoStruct) One(c *gin.Context) {
	var todo models.Todo
	db := utils.DB()

	ID := c.Param("id")

	_, err := todo.Get(db, ID)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error get Todo!"})
		return
	}

	if todo.ID != "" {
		c.JSON(http.StatusOK, todo)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "Todo id not found!"})
}

// Get todo controller
func (t *TodoStruct) Get(c *gin.Context) {
	var todo models.Todo
	db := utils.DB()

	dataTodo, err := todo.Get(db, "")

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error get Todo!"})
		return
	}

	c.JSON(http.StatusOK, dataTodo)
}
