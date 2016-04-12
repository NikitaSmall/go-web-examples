package main

import (
	"encoding/json"

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

// Index runs the index action.
func (c *TasksController) Index(ctx *app.IndexTasksContext) error {
	tasks, _ := json.Marshal(storage.GetTasks())
	return ctx.OK(tasks)
}

// Get runs the get action.
func (c *TasksController) Get(ctx *app.GetTasksContext) error {
	task, _ := storage.GetTask(ctx.ID)
	t, _ := json.Marshal(task)
	return ctx.OK(t)
}

// Create runs the create action.
func (c *TasksController) Create(ctx *app.CreateTasksContext) error {
	storage.AddTask(*ctx.Payload.Text)
	return nil
}

// Delete runs the delete action.
func (c *TasksController) Delete(ctx *app.DeleteTasksContext) error {
	storage.DeleteTask(ctx.ID)
	return nil
}

// Update runs the update action.
func (c *TasksController) Update(ctx *app.UpdateTasksContext) error {
	storage.UpdateTask(ctx.ID, *ctx.Payload.Text, *ctx.Payload.Complete)
	return nil
}
