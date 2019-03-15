package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

// Get 获取标签
func (this *TestController) Get() {
	user := []int{1, 23, 33}
	// data := map[string]int{"张三": 43, "李四": 50}
	// data := map[string]string{
	// 	"code": "200",
	// }
	this.Data["json"] = user
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
