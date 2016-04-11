package routers

import (
	"github.com/astaxie/beego"
	"github.com/nikitasmall/frameworks-comparsion/beego/controllers"
)

func init() {
	beego.Router("/", &controllers.TaskController{}, "get:Index")
	beego.Router("/:id", &controllers.TaskController{}, "get:Get")
	beego.Router("/", &controllers.TaskController{}, "post:Post")
	beego.Router("/:id", &controllers.TaskController{}, "put:Update")
	beego.Router("/:id", &controllers.TaskController{}, "delete:Delete")
}
