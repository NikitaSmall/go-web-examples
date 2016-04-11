package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nikitasmall/frameworks-comparsion/task"
	"strconv"
)

var storage = task.NewStorage()

func indexHandler(c *gin.Context) {
	c.JSON(200, storage.GetTasks())
}

func getTaskHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := storage.GetTask(id)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, task)
	}
}

func addTaskHandler(c *gin.Context) {
	text := c.PostForm("text")
	id := storage.AddTask(text)
	c.JSON(200, gin.H{"id": id})
}

func updateTaskHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	complete, _ := strconv.ParseBool(c.PostForm("complete"))
	task, err := storage.UpdateTask(id, c.PostForm("text"), complete)

	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, task)
	}
}

func deleteTaskHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := storage.DeleteTask(id)

	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, task)
	}
}

func main() {
	r := gin.Default()

	r.GET("/", indexHandler)
	r.GET("/:id", getTaskHandler)

	r.POST("/", addTaskHandler)
	r.PUT("/:id", updateTaskHandler)
	r.DELETE("/:id", deleteTaskHandler)

	r.Run(":3000")
}
