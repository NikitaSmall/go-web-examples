package routers

import (
	"github.com/astaxie/beego"
	"github.com/nikitasmall/frameworks-comparsion/beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.HelloController{})
}
