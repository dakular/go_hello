package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()

	// 注册中间件
	app.Use(myMiddleware)

	// 注册模板
	tmpl := iris.HTML("./views", ".html")
	app.RegisterView(tmpl)

	// 自定义错误页面
	app.OnErrorCode(iris.StatusNotFound, notFound)

	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})

	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./views/hello.html
		ctx.View("index.html")
	})

	app.Handle("GET", "/about", func(ctx iris.Context) {
		ctx.HTML("<h3> I am Duckula </h3>")
	})

	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

	// group route
	api := app.Party("/api", myAuthMiddlewareHandler)

	// methods test
	api.Get("/any", apiHandler)
	api.Get("/any/{username}", apiHandler)
	api.Post("/any", apiHandler)

	app.Run(iris.Addr(":8080"))
}

func notFound(ctx iris.Context) {
	// when 404 then render the template
	// $views_dir/errors/404.html
	ctx.View("errors/404.html")
}

func myMiddleware(ctx iris.Context) {
	// ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

func myAuthMiddlewareHandler(ctx iris.Context) {
	ctx.Application().Logger().Infof("Auth required: %s", ctx.Path())
	ctx.Next()
}

func apiHandler(ctx iris.Context) {
	param_username := ctx.Params().Get("username")
	query_id := ctx.URLParam("id")
	query_datetime := ctx.URLParam("datetime")
	form_username := ctx.FormValue("username")
	if form_username != "" {
		ctx.Values().Set("stored_username", form_username)
		ctx.Application().Logger().Infof("form_username: %s", form_username)
	}
	// query_id := ctx.Request().URL.Query().Get("id")
	// form_username := ctx.FormValueDefault("username", "N/A")
	ctx.Writef("request method: %s\nrequest path: %s\n", ctx.Method(), ctx.Path())
	ctx.Writef("param_username: %s\nquery_id: %s\nquery_datetime: %s\nform_username: %s\n", param_username, query_id, query_datetime, form_username)
	if stored_username := ctx.Values().GetString("stored_username"); stored_username != "" {
		ctx.Writef("stored_username: %s\n", stored_username)
	}
}
