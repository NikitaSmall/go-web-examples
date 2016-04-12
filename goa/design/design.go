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
		Description("Show all tasks")

		Response(OK, "application/json")
	})

	Action("create", func() {
		Routing(POST("/tasks"))
		Description("Create new task")

		Payload(func() {
			Member("text", String, "task text")
		})
	})
})
