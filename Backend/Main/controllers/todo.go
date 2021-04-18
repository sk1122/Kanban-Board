package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"Backend/Main/models"
	_"bytes"
	_"encoding/json"

	"Backend/Main/middleware"
)

type CreateTodoInput struct {
	Todo 	 string `json:"todo"`
	Category string `json:"category"`
}

type UpdateTodoInput struct {
	Todo 	 string `json: "todo"`
	Category string `json: "category"`
}

type UpdateTodoInputList struct {
	Todo  []string `json: "todo"`
	Doing []string `json: "doing"`
	Done  []string `json: "done"`
	Trash []string `json: "trash"`
}

// POST REQUEST
// Creates and Stores Todo's in DB
func CreateTodo(c *gin.Context) {

	User, errr := middleware.GetUser(c.Request)

	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": errr.Error()})
		return
	}

	var input CreateTodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	todo := models.TodoData{Todo: input.Todo, Category: input.Category, Username: User}
	models.DB.Create(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// GET REQUEST
// Fetches All Todo's
func AllTodo(c *gin.Context) {

	User, errr := middleware.GetUser(c.Request)

	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": errr.Error()})
		return
	}

	var todo []models.TodoData

	models.DB.Where("Username = ?", User).Find(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// GET REQUEST
// Fetches Single Todo
func FindTodo(c *gin.Context) {

	User, errr := middleware.GetUser(c.Request)

	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": errr.Error()})
		return
	}

	var todo models.TodoData

	if err := models.DB.Where("id = ? AND Username = ?", c.Param("id"), User).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// PATCH REQUEST
// Updates Required Todo
func UpdateTodo(c *gin.Context) {

	User, errr := middleware.GetUser(c.Request)

	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": errr.Error()})
		return
	}

	var todo models.TodoData

	if err := models.DB.Where("id = ? AND Username = ?", c.Param("id"), User).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateTodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&todo).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodoBool(todo, category, user string) {
	var todoData models.TodoData

	if err := models.DB.Where("Todo = ? AND Username = ?", todo, user).First(&todoData).Error; err != nil {
		return
	}

	input := UpdateTodoInput{Todo: todo, Category: category}
	models.DB.Model(&todoData).Updates(input)
} 

func CreateTodoFromList(todo []string, category, user string) {
	for i:=0;i<len(todo);i++ {
		if middleware.CheckIfTodoExists(todo[i], category, user) {
			continue
		}
		if todo[i] == "" {
			continue
		}
		if middleware.CheckIfTodo(todo[i], user) {
			UpdateTodoBool(todo[i], category, user)
			continue
		}
		todo := models.TodoData{Todo: todo[i], Category: category, Username: user}
		models.DB.Create(&todo)
	}
}

func PostTodoList(c *gin.Context) {

	User, errr := middleware.GetUser(c.Request)

	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": errr.Error()})
		return
	}

	var input UpdateTodoInputList

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return	
	}

	CreateTodoFromList(input.Todo, "Todo", User)
	CreateTodoFromList(input.Done, "Done", User)
	CreateTodoFromList(input.Doing, "Doing", User)
	CreateTodoFromList(input.Trash, "Trash", User)

	c.JSON(http.StatusOK, gin.H{"data": input})
}