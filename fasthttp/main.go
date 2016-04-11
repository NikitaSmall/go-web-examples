package main

// to raw library, we need use more tools fasthttp-* to work with it.

import (
	"fmt"
	"regexp"
	"strconv"

	"encoding/binary"
	"encoding/json"
	"github.com/nikitasmall/frameworks-comparsion/task"
	"github.com/valyala/fasthttp"
)

var storage = task.NewStorage()

var (
	indexRegexp = regexp.MustCompile("/")
	taskRegexp  = regexp.MustCompile("/task/")

	updateTaskRegexp = regexp.MustCompile("/task/update/")
	deleteTaskRegexp = regexp.MustCompile("/task/delete/")
)

// request handler in fasthttp style, i.e. just plain function.
func requestHandler(ctx *fasthttp.RequestCtx) {
	switch {
	case updateTaskRegexp.Match(ctx.Path()):
		if ctx.IsPost() {
			updateTaskHandler(ctx)
		}
	case deleteTaskRegexp.Match(ctx.Path()):
		if ctx.IsPost() {
			deleteTaskHandler(ctx)
		}
	case taskRegexp.Match(ctx.Path()):
		if ctx.IsGet() {
			getTaskHandler(ctx)
		}
		if ctx.IsPost() {
			addTaskHandler(ctx)
		}
	case indexRegexp.Match(ctx.Path()):
		indexHandler(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

func indexHandler(ctx *fasthttp.RequestCtx) {
	tasks, _ := json.Marshal(storage.GetTasks())
	ctx.SetBody(tasks)
}

func addTaskHandler(ctx *fasthttp.RequestCtx) {
	text := string(ctx.FormValue("text"))
	id := storage.AddTask(text)
	fmt.Fprint(ctx, id)
}

func deleteTaskHandler(ctx *fasthttp.RequestCtx) {
	id := int(binary.BigEndian.Uint64(ctx.Path()[len("/task/delete/"):]))
	task, err := storage.DeleteTask(id)

	if err != nil {
		fmt.Fprint(ctx, err.Error())
	} else {
		t, _ := json.Marshal(task)
		ctx.SetBody(t)
	}
}

func updateTaskHandler(ctx *fasthttp.RequestCtx) {
	id := int(binary.BigEndian.Uint64(ctx.Path()[len("/task/update/"):]))
	text := string(ctx.FormValue("text"))
	complete, _ := strconv.ParseBool(string(ctx.FormValue("complete")))

	task, err := storage.UpdateTask(id, text, complete)

	if err != nil {
		fmt.Fprint(ctx, err.Error())
	} else {
		t, _ := json.Marshal(task)
		ctx.SetBody(t)
	}
}

func getTaskHandler(ctx *fasthttp.RequestCtx) {
	id := int(binary.BigEndian.Uint64(ctx.Path()[len("/task/"):]))
	task, err := storage.GetTask(id)

	if err != nil {
		fmt.Fprint(ctx, err.Error())
	} else {
		t, _ := json.Marshal(task)
		fmt.Fprint(ctx, t)
	}
}

func main() {
	fasthttp.ListenAndServe(":8000", requestHandler)
}
