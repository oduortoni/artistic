package routers

import (
	"artistic/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/*", &controllers.MainController{}, "get:Index")
	beego.Router("/api/users", &controllers.UsersController{}, "get:GetUsers")
}
