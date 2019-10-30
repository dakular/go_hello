package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iris-contrib/middleware/csrf"
	"github.com/kataras/iris"
	"math/rand"
	"strconv"
)

type ResponseBean struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func main() {
	app := iris.Default()

	// 注册中间件
	app.Use(myMiddleware)

	// 注册模板
	tmpl := iris.HTML("./views", ".html")
	app.RegisterView(tmpl)

	// 自定义错误页面
	app.OnErrorCode(iris.StatusNotFound, notFound)

	// csrf设置
	// 请注意，提供的身份验证密钥应为32个字节应用程序重新启动时保持不变
	protect := csrf.Protect(
		[]byte("9AB0F421E53A477C084477AEA06096F5"),
		csrf.Secure(false), // Defaults to true, but pass `false` while no https (devmode).
		csrf.Path("/"),     // Path sets the cookie path 设置成根路径`/`否则只能在相同父路径下才有效
	)

	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})

	app.Get("/", protect, homeHandler)
	app.Get("/test/form", protect, formHandler)
	app.Handle("GET", "/about", func(ctx iris.Context) {
		ctx.HTML("<h3> I am Duckula </h3>")
	})

	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

	// group route
	api := app.Party("/api", protect) // ,myAuthMiddlewareHandler
	api.Get("/any", apiHandler)
	api.Get("/any/{username}", apiHandler)
	api.Post("/any", apiHandler)
	api.Post("/post", postApiHandler)
	api.Get("/json", apiHandlerJson)
	api.Get("/sql", apiHandlerSql)

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

func homeHandler(ctx iris.Context) {
	// Bind: {{.message}} with "Hello world!"
	ctx.ViewData("message", "Hello! I am Iris")

	// csrf.TemplateField将CSRF令牌注入即可！视图中只需要一个{{.csrfField}}模板标记
	token := csrf.Token(ctx)
	token_field := csrf.TemplateField(ctx)
	fmt.Println("token", token)
	fmt.Println(csrf.TemplateTag, token_field)
	ctx.ViewData("token", token)
	ctx.ViewData(csrf.TemplateTag, token_field)

	// Render template file: ./views/hello.html
	ctx.View("index.html")
}

func formHandler(ctx iris.Context) {
	token := csrf.Token(ctx)
	token_field := csrf.TemplateField(ctx)
	fmt.Println("token", token)
	fmt.Println(csrf.TemplateTag, token_field)
	ctx.ViewData("token", token)
	ctx.ViewData(csrf.TemplateTag, token_field)
	// ctx.SetCookieKV("csrftoken", token)
	// ctx.SetCookieKV("X-CSRF-Token", token)
	ctx.View("form.html")
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

func postApiHandler(ctx iris.Context) {
	// ctx.Writef("POST OK %s", ctx.FormValue("username"))
	ctx.JSON(&ResponseBean{200, "success", ctx.FormValue("username")})
}

func apiHandlerJson(ctx iris.Context) {
	ctx.JSON(&ResponseBean{200, "success", iris.Map{
		"id":   123,
		"name": "Tiger Liu",
	}})
}

func apiHandlerSql(ctx iris.Context) {
	db, err := sql.Open("mysql", "usr:pwd@tcp(host)/db?charset=utf8")

	if err != nil {
		fmt.Println(err)
		ctx.Writef(err.Error() + "\n")
		return
	}

	// 关闭数据库，db会被多个goroutine共享，可以不调用
	defer db.Close()

	// ping服务器
	if err := db.Ping(); err != nil {
		fmt.Println("opon database fail")
		ctx.Writef(err.Error() + "\n")
		return
	}
	fmt.Println("connnect success")
	ctx.Writef("connnect success" + "\n")

	// 查询数据，指定字段名，返回sql.Rows结果集
	rows, err := db.Query("SELECT id,name from contacts")
	if err != nil {
		fmt.Println("查询出错了")
	}
	// 第一步：接收在数据库表查询到的字段名，返回的是一个string数组切片
	columns, _ := rows.Columns() // columns:  [user_id user_name user_age user_sex]
	fmt.Println(columns)
	// 根据string数组切片的长度构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// 循环读取结果
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
		/*
			// 将每一行的结果都赋值到变量中
			err := rows.Scan(scanArgs...)
			if err != nil {
				fmt.Println(err.Error())
			}
			//将行数据保存到record字典
			record := make(map[string]string)
			for i, col := range values {
				if col != nil {
					//字段名 = 字段信息
					record[columns[i]] = string(col.([]byte))
				}
			}
			fmt.Println(record, record["id"], record["name"])
		*/
	}

	// 查询数据，取所有字段
	rows2, _ := db.Query("select * from contacts")
	// 返回所有列
	cols, _ := rows2.Columns()
	fmt.Println(cols)
	// 这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	// 这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	// 这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	// i := 0
	// result := make(map[int]map[string]string)
	for rows2.Next() {
		//填充数据
		rows2.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		fmt.Println(row)
		_json, _ := json.Marshal(row)
		ctx.Writef(string(_json) + "\n")

		//放入结果集
		// result[i] = row
		//i++
	}
	// fmt.Println(result)

	//准备更新操作
	stmt, err := db.Prepare("UPDATE contacts SET tel=? WHERE id=?")
	if err != nil {
		fmt.Println(err.Error())
	}
	//执行更新操作
	tel := fmt.Sprintf("130%08s", strconv.Itoa(rand.Intn(99999999)))
	fmt.Println("tel", tel)
	res, err := stmt.Exec(tel, 2)
	if err != nil {
		fmt.Println(err.Error())
	}
	//查询更新多少条信息
	num, _ := res.RowsAffected()
	fmt.Println(strconv.FormatInt(num, 10) + " records updated")
	ctx.Writef(strconv.FormatInt(num, 10) + " records updated" + "\n")

}
