package routers

import (
	"simple-api-beego/controllers"
	"simple-api-beego/helpers"
	"simple-api-beego/middlewares"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	// beego.Router("/api/v1/", &controllers.MainController{})

	web.InsertFilter("/api/v1", web.BeforeRouter, helpers.ForbiddenHandler)

	auth := web.NewNamespace("/api/v1/auth",
		web.NSRouter("/login", &controllers.AuthController{}, "post:Login"),
		web.NSRouter("/register", &controllers.UserController{}, "post:Create"),
	)

	users := web.NewNamespace("/api/v1/users",
		web.NSBefore(middlewares.JWTMiddleware),
		web.NSRouter("", &controllers.UserController{}, "get:GetAll"),
		web.NSRouter("/:id", &controllers.UserController{}, "get:GetUserByID;put:Update;delete:Delete"),
	)

	web.AddNamespace(auth, users)
}
