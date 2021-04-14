package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"Backend/Main/models"
	"fmt"

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

	fmt.Println(todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func AllTodoList(c *gin.Context) {

	User, errr := middleware.GetUser(c.Request)

	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": errr.Error()})
		return
	}

	var todo models.TodoDataList

	models.DB.Where("Username = ?", User).First(&todo)

	fmt.Println(todo)

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

func PostTodoList(c *gin.Context) {

	User, errr := middleware.GetUser(c.Request)

	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": errr.Error()})
		return
	}

	var input models.TodoDataList

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list := models.TodoDataList{Todo: input.Todo, Doing: input.Doing, Done: input.Done, Trash: input.Trash, Username: User}

	models.DB.Create(&list)

	c.JSON(http.StatusOK, gin.H{"data": list})
}