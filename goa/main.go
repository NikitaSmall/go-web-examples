package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/nikitasmall/frameworks-comparsion/goa/app"
	"github.com/nikitasmall/frameworks-comparsion/goa/swagger"
)

func main() {
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "tasks" controller
	c := NewTasksController(service)
	app.MountTasksController(service, c)
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
