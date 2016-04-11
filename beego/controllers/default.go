package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/nikitasmall/frameworks-comparsion/beego/models"
)

type TaskController struct {
	beego.Controller
}

func (c *TaskController) Index() {
	c.Data["json"] = models.Storage.GetTasks()
	c.ServeJSON()
}

func (c *TaskController) Get() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param("id"))
	c.Data["json"], _ = models.Storage.GetTask(id)
	c.ServeJSON()
}

func (c *TaskController) Post() {
	text := c.GetString("text")
	c.Data["json"] = models.Storage.AddTask(text)
	c.ServeJSON()
}

func (c *TaskController) Update() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param("id"))
	text := c.GetString("text")
	complete, _ := c.GetBool("complete")

	c.Data["json"], _ = models.Storage.UpdateTask(id, text, complete)
	c.ServeJSON()
}

func (c *TaskController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param("id"))
	c.Data["json"], _ = models.Storage.DeleteTask(id)
	c.ServeJSON()
}
