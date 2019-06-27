package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//response JSON
type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[string]*Todo{}

func getTodosHandler(c *gin.Context) {
	todo := []*Todo{}
	for _, td := range todos {
		todo = append(todo, td)
	}
	c.JSON(http.StatusOK, todo)
}

func getTodoByIDHandler(c *gin.Context) {
	id := c.Param("id")
	todo, result := todos[id]
	if !result {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func createTodosHandler(c *gin.Context) {
	todo := Todo{}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	i := len(todos)
	i++
	id := strconv.Itoa(i)
	todo.ID = id
	todos[id] = &todo
	c.JSON(http.StatusCreated, todo)

}

func updateTodosHandler(c *gin.Context) {
	id := c.Param("id")
	todo := todos[id]
	if err := c.ShouldBindJSON(todo); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)

}

func deleteTodosHandler(c *gin.Context) {
	id := c.Param("id")
	delete(todos, id)
	c.JSON(http.StatusOK, gin.H{"status": "success"})

}

func main() {
	router := gin.Default()
	router.GET("/api/todos", getTodosHandler)
	router.GET("/api/todos/:id", getTodoByIDHandler)
	router.POST("/api/todos", createTodosHandler)
	router.PUT("/api/todos/:id", updateTodosHandler)
	router.DELETE("/api/todos/:id", deleteTodosHandler)
	router.Run(":1234")
}
