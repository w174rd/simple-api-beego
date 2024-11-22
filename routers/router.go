package routers

import (
	"simple-api-beego/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/api/v1/", &controllers.MainController{})
	beego.Router("/api/v1/users", &controllers.UserController{}, "get:GetAll")
	beego.Router("/api/v1/users", &controllers.UserController{}, "post:Create")
	beego.Router("/api/v1/users/:id", &controllers.UserController{}, "put:Update")
	beego.Router("/api/v1/users/:id", &controllers.UserController{}, "delete:Delete")
}
