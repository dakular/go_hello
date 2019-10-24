package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"hello/test_beego/controllers"
	"strconv"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/user/", &controllers.UserController{}) // "*:Get"
	beego.Router("/api/user/list", &controllers.UserController{}, "get:ListUser")
	beego.Router("/api/user/:id", &controllers.UserController{}, "get:ShowUser")
	beego.Router("/api/user/create", &controllers.UserController{}, "post:CreateUser")

	// 简单路由
	beego.Get("/api/get", func(ctx *context.Context) {
		ctx.Output.Body([]byte("This is GET method"))
	})
	beego.Post("/api/post", func(ctx *context.Context) {
		ctx.Output.Body([]byte("This is POST method, username=" + ctx.Input.Query("username")))
	})
	beego.Any("/api/any", func(ctx *context.Context) {
		ctx.Output.Body([]byte("This router can accept any method\n" + ctx.Input.IP() + ":" + strconv.Itoa(ctx.Input.Port()) + "\n" + ctx.Input.URL() + "\n" + ctx.Input.Method() + "\nname: " + ctx.Input.Query("name")))
	})

	// 注册注解路由
	beego.Include(&controllers.CMSController{})
}
