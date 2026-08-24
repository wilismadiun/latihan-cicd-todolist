package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddTodoList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		Todos = make(map[string]Todo)

		router := gin.New()
		router.POST("/todos", AddTodoList)

		body := `{
			"name": "Belajar Go",
			"completed": false
		}`

		req := httptest.NewRequest(
			http.MethodPost,
			"/todos",
			bytes.NewBufferString(body),
		)

		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		var response map[string]interface{}

		err := json.Unmarshal(rec.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.Equal(t, "data berhasil ditambahkan", response["message"])

		data := response["data"].(map[string]interface{})

		assert.NotEmpty(t, data["id"])
		assert.Equal(t, "Belajar Go", data["name"])
		assert.Equal(t, false, data["completed"])

		assert.Len(t, Todos, 1)
	})
}

func TestAddTodoList_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	Todos = make(map[string]Todo)

	router := gin.New()
	router.POST("/todos", AddTodoList)

	body := `{
		"name": "Belajar Go",
		"completed":
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/todos",
		bytes.NewBufferString(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}

	err := json.Unmarshal(rec.Body.Bytes(), &response)

	assert.NoError(t, err)

	assert.Equal(
		t,
		"gagal menambahkan todo list",
		response["message"],
	)

	assert.Len(t, Todos, 0)
}

func TestGetTodoList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	Todos = map[string]Todo{
		"todo-1": {
			Name:      "Belajar Go",
			Completed: false,
		},
		"todo-2": {
			Name:      "Belajar Docker",
			Completed: true,
		},
	}

	router := gin.New()
	router.GET("/todos", GetTodoList)

	req := httptest.NewRequest(
		http.MethodGet,
		"/todos",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response struct {
		Message string `json:"message"`
		Data    []struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			Completed bool   `json:"completed"`
		} `json:"data"`
	}

	err := json.Unmarshal(rec.Body.Bytes(), &response)

	assert.NoError(t, err)

	assert.Equal(
		t,
		"data berhasil ditampilkan",
		response.Message,
	)

	assert.Len(t, response.Data, 2)

	assert.NotEmpty(t, response.Data[0].ID)
	assert.NotEmpty(t, response.Data[0].Name)
}

func TestGetTodoList_Empty(t *testing.T) {
	gin.SetMode(gin.TestMode)

	Todos = make(map[string]Todo)

	router := gin.New()
	router.GET("/todos", GetTodoList)

	req := httptest.NewRequest(
		http.MethodGet,
		"/todos",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response struct {
		Message string `json:"message"`
		Data    []struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			Completed bool   `json:"completed"`
		} `json:"data"`
	}

	err := json.Unmarshal(rec.Body.Bytes(), &response)

	assert.NoError(t, err)

	assert.Equal(
		t,
		"data berhasil ditampilkan",
		response.Message,
	)

	assert.Empty(t, response.Data)
}

func TestGetTodoByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	Todos = map[string]Todo{
		"todo-1": {
			Name:      "Belajar Go",
			Completed: false,
		},
	}

	router := gin.New()
	router.GET("/todos/:id", GetTodoByID)

	req := httptest.NewRequest(
		http.MethodGet,
		"/todos/todo-1",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response struct {
		Message string `json:"message"`
		Data    struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			Completed bool   `json:"completed"`
		} `json:"data"`
	}

	err := json.Unmarshal(rec.Body.Bytes(), &response)

	assert.NoError(t, err)

	assert.Equal(
		t,
		"data berhasil ditampilkan",
		response.Message,
	)

	assert.Equal(t, "todo-1", response.Data.ID)
	assert.Equal(t, "Belajar Go", response.Data.Name)
	assert.Equal(t, false, response.Data.Completed)
}

func TestGetTodoByID_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	Todos = map[string]Todo{
		"todo-1": {
			Name:      "Belajar Go",
			Completed: false,
		},
	}

	router := gin.New()
	router.GET("/todos/:id", GetTodoByID)

	req := httptest.NewRequest(
		http.MethodGet,
		"/todos/todo-999",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	var response map[string]interface{}

	err := json.Unmarshal(rec.Body.Bytes(), &response)

	assert.NoError(t, err)

	assert.Equal(
		t,
		"todo tidak ditemukan",
		response["message"],
	)
}

func TestDeleteTodo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	Todos = map[string]Todo{
		"todo-1": {
			Name:      "Belajar Go",
			Completed: false,
		},
	}

	router := gin.New()
	router.DELETE("/todos/:id", DeleteTodo)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/todos/todo-1",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}

	err := json.Unmarshal(rec.Body.Bytes(), &response)

	assert.NoError(t, err)

	assert.Equal(
		t,
		"todo berhasil dihapus",
		response["message"],
	)

	_, exists := Todos["todo-1"]

	assert.False(t, exists)
	assert.Len(t, Todos, 0)
}
