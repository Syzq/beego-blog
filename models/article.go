package models

import (
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Desc       string    `json:"desc"`
	Content    string    `json:"content"`
	CreatedOn  time.Time `json:"created_on" orm:"auto_now_add;type(datetime)"` //auto_now_add 第一次保存时才设置时间
	CreatedBy  string    `json:"created_by" orm:"size(60)"`
	ModifiedOn time.Time `json:"modified_on" orm:"auto_now;type(datetime)"` //auto_now 每次model 保存时都会对时间自动更新
	ModifiedBy string    `json:"modified_by" orm:"size(60)"`

	Tags []*Tag `orm:"rel(m2m)"`
}

func init() {
	orm.RegisterModel(new(Article))
	// insertTestData()
}

func AddArticle(article Article) bool {

	o := orm.NewOrm()

	_, err := o.Insert(&article)
	if err != nil {
		return false
	}
	return true
}

// InsertTestData 多对多插入测试
func InsertTestData() {
	o := orm.NewOrm()

	var article1 Article
	article1.Desc = "多对多描述"
	o.Insert(&article1)

	var tag1 Tag
	tag1.Name = "先插入文章，再插入标签"
	o.Insert(&tag1)

	m2m := o.QueryM2M(&article1, "Tags")
	m2m.Add(&tag1)
	beego.Info("duo dui duo")
}

// ReadM2M 多对多查询
func ReadM2M() (article []Article) {
	o := orm.NewOrm()
	// var article Article
	if err := o.QueryTable("article").Filter("Id",5).RelatedSel().One(&article); err != nil {
		beego.Error(err)
		if _, err := o.LoadRelated(&article, "Tags"); err != nil {
			beego.Error(err)
		}
	}
	beego.Info("多对多查询", article)
	return article
	// var tag Tag
	// o.QueryTable("tag").Filter("Id", 6).RelatedSel().One(&tag)
	// o.LoadRelated(&tag, "Articles")
	// beego.Info("多对多查询", tag)
}
