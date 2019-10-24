package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type JSONStruct struct {
	Code int
	Data string
}

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	list_url := beego.URLFor("UserController.ListUser")
	show_url := beego.URLFor("UserController.ShowUser", ":id", "userid")
	create_url := beego.URLFor("UserController.CreateUser")
	c.Ctx.WriteString("defaut get method\n" + list_url + "\n" + show_url + "\n" + create_url)
}

func (c *UserController) ShowUser() {
	id := c.Ctx.Input.Param(":id")
	println(id)
	c.Data["ID"] = id
	// c.TplName = "index.tpl"
	// c.Ctx.Output.Body([]byte(id))
	c.Ctx.WriteString("SHOW USER, ID=" + id)
}

func (c *UserController) ListUser() {
	// c.Ctx.WriteString("User List...")
	c.Data["not_used_param"] = "123"
	user_list := []string{"user1", "user2", "user3", "user4", "user5"}
	mystruct := &JSONStruct{0, strings.Join(user_list[:], ", ")}
	c.Data["json"] = mystruct
	// c.Data["json"] = "{\"UserList\":\"" + strings.Join(user_list[:], ", ") + "\"}"
	c.ServeJSON()
}

func (c *UserController) CreateUser() {
	c.Ctx.WriteString("[POST] Create User... username=" + c.Ctx.Input.Query("username"))
}
