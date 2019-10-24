package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["hello/test_beego/controllers:CMSController"] = append(beego.GlobalControllerRouter["hello/test_beego/controllers:CMSController"],
        beego.ControllerComments{
            Method: "ShowCMS",
            Router: `/api/cms/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello/test_beego/controllers:CMSController"] = append(beego.GlobalControllerRouter["hello/test_beego/controllers:CMSController"],
        beego.ControllerComments{
            Method: "ListCMS",
            Router: `/api/cms/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
