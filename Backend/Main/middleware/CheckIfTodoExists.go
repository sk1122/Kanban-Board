package middleware

import (
	"Backend/Main/models"
)

func CheckIfTodoExists(todo, category, user string) bool {
	var todoData models.TodoData
	if err := models.DB.Where("Todo = ? AND Category = ? AND Username = ?", todo, category, user).First(&todoData).Error; err != nil {
		return false
	}
	return true
}


func CheckIfTodo(todo, user string) bool {
	var todoData models.TodoData

	if err := models.DB.Where("Todo = ? AND Username = ?", todo, user).First(&todoData).Error; err != nil {
		return false
	}

	return true
}