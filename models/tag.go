package models

import (
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

// Tag model
type Tag struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	CreatedOn  time.Time `json:"created_on" orm:"auto_now_add;type(datetime)"` //auto_now_add 第一次保存时才设置时间
	CreatedBy  string    `json:"created_by" orm:"size(60)"`
	ModifiedOn time.Time `json:"modified_on" orm:"auto_now;type(datetime)"` //auto_now 每次model 保存时都会对时间自动更新
	ModifiedBy string    `json:"modified_by" orm:"size(60)"`
	State      int       `json:"state" orm:"size(4);default(1)"` //状态：0 为禁用，1为启用
}

func init() {
	orm.RegisterModel(new(Tag))
}

// AddTag 创建标签
func AddTag(tag *Tag) bool {
	o := orm.NewOrm()

	o.Insert(tag)

	return true
}

func GetTags() {

}

func GetTagTotal() (count int) {
	o := orm.NewOrm()

	cnt, err := o.QueryTable("tag").Count()
	beego.Error("Count Num: %s, %s", cnt, err)
	return
}
