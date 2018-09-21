package models;

import (
	"github.com/astaxie/beego/orm";

	// "github.com/astaxie/beego"
);

// Model Struct
type User struct { // 表名将会自动转换；如：User-> user, DB_AuthUser -> d_b__auth_user ； 除了开头的大写字母以外，遇到大写会增加_，原名称中的下划线保留。
	Id int // 同样字段名也是头一个字母会转为小写
	// Id int `orm:"pk;auto"` // orm: 为声明建表时字段类型，默认主键为Id，可以不写pk和auto
	Name string "orm:\"size(100)\"" // ``这个符号和""是一样的，只是为了在里面能使用双引号或多行时会这样写
	Email string `orm:"size(200);default()"` // default():默认值为空字符串(可以省略不写)，如果写成：default(NULL)就成默认字符为：“NULL”字符串；貌似在程序中字符串不能设置为默认null
}

// func (u *User) TableName() string { // 自定义表名（如果表名和结构体名不一致时）
//     return "auth_user";
// }

func init() { // init() 方法，在执行完本脚本的常量和全局变量后，会自动调用此方法
	orm.RegisterModel(new(User)); // register model  .   new(Struct):创建一个新的结构体对象
}
