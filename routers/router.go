package routers

import (
	"github.com/astaxie/beego"
	"github.com/wms-go/beeo-demo/controllers"
)

func init() {
   // beego.Router("/", &controllers.MainController{})
   beego.Include(&controllers.UserController{})
}
