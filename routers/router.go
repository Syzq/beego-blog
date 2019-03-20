package routers

import (
	"beego-blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/tag",
			beego.NSInclude(
				&controllers.TagController{},
			),
		),
		beego.NSNamespace("/article",
			beego.NSInclude(
				&controllers.ArticleController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
