package main

type Todo struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var Todos = make(map[string]Todo)
