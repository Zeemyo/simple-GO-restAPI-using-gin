package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ini struktur penting untuk pengelolaan data
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Cleanng Room", Completed: false},
	{ID: "2", Item: "Gaming", Completed: false},
	{ID: "3", Item: "Recording", Completed: false},
	{ID: "4", Item: "Working", Completed: false},
}

// mengambil data array todo dan di ubah ke JSON
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	//mengecek apakah format struct todo ada json nya atau tidak, jika tidak maka akan lanjut
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	//kesini, disini menggabungkan variable array todos dengan variable newTodo sebagai JSON
	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo Not Found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
