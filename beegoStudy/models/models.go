package models;

import (
	"github.com/astaxie/beego/orm";
);

type Member struct { // 表member只有三个字段：id，name，profile_id 
	Id       int
	Name     string
	Profile  *Profile    `orm:"rel(one)"` // One to one relation 
    Article  []*Article  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct { // 表profile只有两个字段：id，age
	Id       int
	Age      int16
	Member   *Member  `orm:"reverse(one)"` // 设置一对一的反向关系（可选）
}

type Article struct { // 表有id，title，member_id三个字段
	Id       int
	Title    string
	Member   *Member  `orm:"rel(fk)"` // 设置一对多关系
	Tags     []*Tag   `orm:"rel(m2m)"` // 设置多对多关系。 多对多的过程中，一定需要另外一张表来保存他们的关系
	// 这里会多出一张表；表article_tags的字段为：id，article_id,tag_id
}

type Tag struct {
	Id       int
	Name     string
	// Articles []*Article `orm:"reverse(many)"` // reverse 都是可以省略的
}

func init() {
	orm.RegisterModel(new(Member), new(Profile), new(Article), new(Tag)); // 注册多个模型
	// orm.RegisterModelWithPrefix("prefix_", new(Member));  //  使用表前缀，创建后的表名为：prefix_member
}
