package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego-blog/controllers:TagController"] = append(beego.GlobalControllerRouter["beego-blog/controllers:TagController"],
        beego.ControllerComments{
            Method: "GetTags",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-blog/controllers:TagController"] = append(beego.GlobalControllerRouter["beego-blog/controllers:TagController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/\:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-blog/controllers:TagController"] = append(beego.GlobalControllerRouter["beego-blog/controllers:TagController"],
        beego.ControllerComments{
            Method: "GetAllTags",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
