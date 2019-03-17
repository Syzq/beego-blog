package routers

import (
	"beego-blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/tag", &controllers.TagController{})
}
