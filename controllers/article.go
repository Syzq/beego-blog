package controllers

import (
	"beego-blog/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

// @router / [post]
func (c *ArticleController) AddArticle() {
	var article models.Article
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &article); err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	res := models.AddArticle(article)
	c.Data["json"] = res
	c.ServeJSON()
}
