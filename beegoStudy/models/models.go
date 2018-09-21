package models;

import (
	"github.com/astaxie/beego/orm";
);

type Member struct {
	Id       int
	Name     string
	Profile  *Profile    `orm:"rel(one)"` // One to one relation
    Article  []*Article  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id       int
	Age      int16
	Member   *Member  `orm:"reverse(one)"` // 设置一对一的反向关系（可选）
}

type Article struct {
	Id       int
	Title    string
	Member   *Member  `orm:"rel(fk)"` // 设置一对多关系
	Tags     []*Tag   `orm:"rel(m2m)"` // 设置多对多关系
}

type Tag struct {
	Id       int
	Name     string
	Articles []*Article `orm:"reverse(many)"`
}

func init() {
    orm.RegisterModel(new(Member), new(Profile), new(Article), new(Tag)); // 注册多个模型
}
