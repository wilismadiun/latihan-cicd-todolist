package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.POST("/todos", AddTodoList)
	router.GET("/todos", GetTodoList)
	router.GET("/todos/:id", GetTodoByID)
	router.DELETE("/todos/:id", DeleteTodo)

	router.Run(":3000")
}
