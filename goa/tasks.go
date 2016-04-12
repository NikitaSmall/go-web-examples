package main

import (
	"encoding/json"
	"fmt"

	"github.com/goadesign/goa"

	"github.com/nikitasmall/frameworks-comparsion/goa/app"
	"github.com/nikitasmall/frameworks-comparsion/task"
)

var storage = task.NewStorage()

// TasksController implements the tasks resource.
type TasksController struct {
	*goa.Controller
}

// NewTasksController creates a tasks controller.
func NewTasksController(service *goa.Service) *TasksController {
	return &TasksController{Controller: service.NewController("TasksController")}
}

// Create runs the create action.
func (c *TasksController) Create(ctx *app.CreateTasksContext) error {
	fmt.Print(ctx.FormValue("text"))
	storage.AddTask(ctx.FormValue("text"))
	return nil
}

// Index runs the index action.
func (c *TasksController) Index(ctx *app.IndexTasksContext) error {
	tasks, _ := json.Marshal(storage.GetTasks())
	return ctx.OK(tasks)
}
