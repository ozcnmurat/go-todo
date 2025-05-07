package main

import (
	"todo-app/handlers"   // projenizin handlers paketi
	"todo-app/middleware" // projenizin middleware paketi

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Auth routes
	router.POST("/login", handlers.Login)

	// Protected routes
	authorized := router.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	{
		// TODO List routes
		authorized.GET("/todos", handlers.GetTodos)
		authorized.POST("/todos", handlers.CreateTodo)
		authorized.PUT("/todos/:id", handlers.UpdateTodo)
		authorized.DELETE("/todos/:id", handlers.DeleteTodo)

		// TODO Item routes
		authorized.POST("/todos/:id/items", handlers.AddTodoItem)
		authorized.PUT("/todos/:id/items/:itemId", handlers.UpdateTodoItem)
		authorized.DELETE("/todos/:id/items/:itemId", handlers.DeleteTodoItem)
	}

	router.Run(":8080")
}
