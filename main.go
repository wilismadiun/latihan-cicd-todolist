package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	fmt.Println("Build from commit: $CODEBUILD_RESOLVED_SOURCE_VERSION")

	router.POST("/todos", AddTodoList)
	router.GET("/todos", GetTodoList)
	router.GET("/todos/:id", GetTodoByID)
	router.DELETE("/todos/:id", DeleteTodo)

	router.Run(":3000")
}
