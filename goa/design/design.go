package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("goa", func() {
	Title("Tasks API")
	Description("Example API for tasks")
	Host("localhost:3000")
	Scheme("http")
})

var _ = Resource("tasks", func() {
	Action("index", func() {
		Routing(GET("/tasks"))
		Description("Shows all tasks")

		Response(OK, "application/json")
	})

	Action("get", func() {
		Routing(GET("/tasks/:id"))
		Description("Shows the task")

		Params(func() {
			Param("id", Integer, "ID of the task")
		})

		Response(OK, "application/json")
	})

	Action("create", func() {
		Routing(POST("/tasks"))
		Description("Creates new task")

		Payload(func() {
			Member("text", String, "task text")
		})
	})

	Action("update", func() {
		Routing(PUT("/tasks/:id"))
		Description("Updates some task")

		Params(func() {
			Param("id", Integer, "ID of the task")
		})

		Payload(func() {
			Member("text", String, "task text")
			Member("complete", Boolean, "task completion")
		})
	})

	Action("delete", func() {
		Routing(DELETE("/tasks/:id"))
		Description("Deletes some task")

		Params(func() {
			Param("id", Integer, "ID of the task")
		})
	})
})
