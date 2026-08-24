package main

import (
	"net/http"
	"todo-app/generate"

	"github.com/gin-gonic/gin"
)

func AddTodoList(c *gin.Context) {
	var todo Todo

	err := c.ShouldBindBodyWithJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "gagal menambahkan todo list",
			"error":   err.Error(),
		})
		return
	}

	id := generate.GenerateID()

	Todos[id] = todo

	c.JSON(http.StatusCreated, gin.H{
		"message": "data berhasil ditambahkan",
		"data": gin.H{
			"id":        id,
			"name":      todo.Name,
			"completed": todo.Completed,
		},
	})
}

func GetTodoList(c *gin.Context) {
	data := make([]gin.H, 0, len(Todos))

	for id, todo := range Todos {
		data = append(data, gin.H{
			"id":        id,
			"name":      todo.Name,
			"completed": todo.Completed,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "data berhasil ditampilkan",
		"data":    data,
	})
}

func GetTodoByID(c *gin.Context) {
	id := c.Param("id")

	todo, exists := Todos[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "todo tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "data berhasil ditampilkan",
		"data": gin.H{
			"id":        id,
			"name":      todo.Name,
			"completed": todo.Completed,
		},
	})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	_, exists := Todos[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "todo tidak ditemukan",
		})
		return
	}

	delete(Todos, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "todo berhasil dihapus",
	})
}
