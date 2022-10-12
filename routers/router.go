package routers

import (
	"gmis/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/_nav", &controllers.NavController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/_page", &controllers.PageController{})
}
