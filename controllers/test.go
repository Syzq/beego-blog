package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

// Get 获取标签
func (this *TestController) Get() {

	name := this.Ctx.Input.Query("name")
	age := this.Ctx.Input.Query("age")

	data := make(map[string]interface{})
	data["name"] = name
	data["age"] = age

	this.Data["json"] = data
	this.ServeJSON()
}

// Post 创建标签
func (c *TestController) Post() {
	beego.Error("post")
}

// Put 修改标签
func (c *TestController) Put() {
	beego.Error("put")
}

// Delete 删除标签
func (c *TestController) Delete() {
	beego.Error("delete")
}
