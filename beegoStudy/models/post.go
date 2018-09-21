package models;

import (
	"github.com/astaxie/beego/orm";
);

type Post struct { // model struct
	Id     int     `orm:"auto"`; // int4 自增主键
	Title  string  `orm:"size(100)"`; // varchar(100) default ''
	User   *User   `orm:"rel(fk)"`; // "user_id" int4 NOT NULL; 查询当前表时，会自动进行连表，连表的依据是：ON user."id" = post."user_id"
	// User 查询出来的数据是一个User结构体
}

func init() { // 必须先注册才能使用，不然无法使用报：表不存在table name: `post` not exists
	orm.RegisterModel(new(Post)); // register model
}
