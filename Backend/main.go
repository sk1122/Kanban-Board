package main

import (
	"github.com/gin-gonic/gin"

	"Backend/Main/models"
	"Backend/Auth/db"

	"Backend/Main/controllers"
	"Backend/Auth/views"

	"Backend/Main/middleware"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()
	db.ConnectDatabase()

	router.POST("/register", views.RegisterUser)
	router.POST("/login", views.LoginUser)
	router.GET("/demo", views.Demo)

	router.GET("/get", middleware.TokenAuthMiddleware(), controllers.AllTodo)
	router.POST("/todo",middleware.TokenAuthMiddleware(), controllers.CreateTodo)
	router.GET("/todo/:id",middleware.TokenAuthMiddleware(), controllers.FindTodo)
	router.PATCH("/todo/update/:id",middleware.TokenAuthMiddleware(), controllers.UpdateTodo)
	router.POST("/todo/list",middleware.TokenAuthMiddleware(), controllers.PostTodoList)
	router.GET("/list",middleware.TokenAuthMiddleware(), controllers.AllTodoList)

	router.Run()
}