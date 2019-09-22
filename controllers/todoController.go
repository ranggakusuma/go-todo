package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranggakusuma/go-todo/models"
	"github.com/ranggakusuma/go-todo/utils"
)

// TodoStruct struct
type TodoStruct struct{}

var todoController TodoStruct

// Routes function for todos
func (t *TodoStruct) Routes(route *gin.Engine) {
	todoRoute := route.Group("api/todos")

	todoRoute.GET("/", todoController.Get)
	todoRoute.GET("/:id", todoController.One)
	todoRoute.POST("/", todoController.Create)
	todoRoute.DELETE("/:id", todoController.Delete)
}

// Create todo controller
func (t *TodoStruct) Create(c *gin.Context) {
	var todo models.Todo
	db := utils.DB()
	tx := db.Begin()
	// tx := db
	c.BindJSON(&todo)

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

// Delete controller
func (t *TodoStruct) Delete(c *gin.Context) {
	var todo models.Todo
	db := utils.DB()

	ID := c.Param("id")
	tx := db.Begin()

	err := todo.Delete(tx, ID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tx.Commit()

	c.JSON(http.StatusAccepted, gin.H{"message": "Todo item deleted successfully!"})
}
