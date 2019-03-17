package controllers

import (
	"beego-blog/models"
	"beego-blog/pkg/e"
	"encoding/json"

	"github.com/astaxie/beego/validation"

	"github.com/astaxie/beego"
)

type TagController struct {
	beego.Controller
}

// Get 获取标签
func (c *TagController) Get() {
	beego.Error("get")

	c.Data["json"] = models.GetTagTotal()
	c.ServeJSON()
}

// Post 创建标签
func (c *TagController) Post() {
	returnData := make(map[string]interface{})
	// 获取 RequestBody 中的值
	var tag models.Tag
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &tag); err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	// 表单验证
	valid := validation.Validation{}
	valid.Required(tag.Name, "name").Message("名称不能为空")
	valid.MaxSize(tag.Name, 20, "name").Message("名称太长了")
	valid.Range(tag.State, 0, 1, "state").Message("状态只允许 0 或 1")

	if !valid.HasErrors() {
		models.AddTag(&tag)
		returnData["code"] = e.SUCCESS
		returnData["msg"] = e.GetMsg(e.SUCCESS)
		// returnData["data"] = make(map[string]string)
		c.Data["json"] = returnData
	} else {
		for _, err := range valid.Errors {
			returnData["code"] = e.INVALID_PARAMS
			returnData["msg"] = err.Message
			c.Data["json"] = returnData
		}
	}

	// 返回 json 数据
	c.ServeJSON()
}

// Put 修改标签
func (c *TagController) Put() {
	beego.Error("put")
}

// Delete 删除标签
func (c *TagController) Delete() {
	beego.Error("delete")
}
