package handlers

import (
	"strconv"
	"todo-app/models"
	"todo-app/services"

	"github.com/gin-gonic/gin"
)

// GetTodos tüm todoları getirir
func GetTodos(c *gin.Context) {
	userID := c.GetUint("user_id")
	userType := c.GetInt("user_type")

	todoService := services.GetTodoService()
	todos := todoService.GetAllTodos(userID, userType == 2)

	c.JSON(200, todos)
}

// GetTodo belirli bir todoyu getirir
func GetTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	todoService := services.GetTodoService()
	todo, err := todoService.GetTodo(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")
	userType := c.GetInt("user_type")
	if userType != 2 && todo.UserID != userID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	c.JSON(200, todo)
}

// CreateTodo yeni bir todo oluşturur
func CreateTodo(c *gin.Context) {
	var todo models.TodoList
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")
	todo.UserID = userID

	todoService := services.GetTodoService()
	if err := todoService.CreateTodo(&todo); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, todo)
}

// UpdateTodo bir todoyu günceller
func UpdateTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var todo models.TodoList
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todoService := services.GetTodoService()
	existing, err := todoService.GetTodo(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")
	userType := c.GetInt("user_type")
	if userType != 2 && existing.UserID != userID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	todo.ID = uint(id)
	todo.UserID = existing.UserID

	if err := todoService.UpdateTodo(&todo); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, todo)
}

// DeleteTodo bir todoyu siler
func DeleteTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	todoService := services.GetTodoService()
	existing, err := todoService.GetTodo(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")
	userType := c.GetInt("user_type")
	if userType != 2 && existing.UserID != userID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	if err := todoService.DeleteTodo(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Todo deleted successfully"})
}

// AddTodoItem bir todoya yeni item ekler
func AddTodoItem(c *gin.Context) {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid todo ID"})
		return
	}

	var item models.TodoItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todoService := services.GetTodoService()
	todo, err := todoService.GetTodo(uint(todoID))
	if err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	userID := c.GetUint("user_id")
	userType := c.GetInt("user_type")
	if userType != 2 && todo.UserID != userID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	if err := todoService.CreateTodoItem(uint(todoID), &item); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, item)
}

// UpdateTodoItem bir todo itemı günceller
func UpdateTodoItem(c *gin.Context) {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid todo ID"})
		return
	}

	itemID, err := strconv.ParseUint(c.Param("itemId"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid item ID"})
		return
	}

	var item models.TodoItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todoService := services.GetTodoService()
	todo, err := todoService.GetTodo(uint(todoID))
	if err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	userID := c.GetUint("user_id")
	userType := c.GetInt("user_type")
	if userType != 2 && todo.UserID != userID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	item.ID = uint(itemID)
	item.TodoListID = uint(todoID)

	if err := todoService.UpdateTodoItem(&item); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, item)
}

// DeleteTodoItem bir todo itemı siler
func DeleteTodoItem(c *gin.Context) {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid todo ID"})
		return
	}

	itemID, err := strconv.ParseUint(c.Param("itemId"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid item ID"})
		return
	}

	todoService := services.GetTodoService()
	todo, err := todoService.GetTodo(uint(todoID))
	if err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	userID := c.GetUint("user_id")
	userType := c.GetInt("user_type")
	if userType != 2 && todo.UserID != userID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	if err := todoService.DeleteTodoItem(uint(todoID), uint(itemID)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Todo item deleted successfully"})
}
