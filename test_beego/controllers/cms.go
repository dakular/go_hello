package controllers

import (
	"github.com/astaxie/beego"
)

type CMSController struct {
	beego.Controller
}

func (c *CMSController) URLMapping() {
	c.Mapping("ListCMS", c.ListCMS)
	c.Mapping("ShowCMS", c.ShowCMS)
}

// @router /api/cms/:id [get]
func (c *CMSController) ShowCMS() {
	id := c.Ctx.Input.Param(":id")
	println(id)
	c.Data["ID"] = id
	// c.TplName = "index.tpl"
	// c.Ctx.Output.Body([]byte(id))
	c.Ctx.WriteString("SHOW CMS, ID=" + id)
}

// @router /api/cms/list [get]
func (c *CMSController) ListCMS() {
	c.Ctx.WriteString("CMS List...")
}

func (c *CMSController) CreateCMS() {
	c.Ctx.WriteString("[POST] Create CMS...")
}
