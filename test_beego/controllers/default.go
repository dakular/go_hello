package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "duckula.net"
	c.Data["Email"] = "i@duckula.net"
	c.Data["DB"] = beego.AppConfig.String("mysqldb")
	c.Data["xsrf_data"] = template.HTML(c.XSRFFormHTML())
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "index.tpl"
}
