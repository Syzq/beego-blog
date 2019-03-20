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
	State      int       `json:"state" orm:"size(10);default(1)"` //状态：0 为禁用，1为启用

	Articles []*Article `orm:"reverse(many)"`
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

func GetTags(pageSize int) (tag []Tag) {
	o := orm.NewOrm()

	// tag := Tag{Name: name}
	// o.Read(&tag, "Name")

	// beego.Info(tag)
	// return tags

	qs := o.QueryTable("tag")
	qs.Limit(pageSize).All(&tag)
	return tag
}

func GetTagTotal() (count int64) {
	o := orm.NewOrm()

	cnt, err := o.QueryTable("tag").Count()
	beego.Info("Count Num:", cnt, err)
	return cnt
}
