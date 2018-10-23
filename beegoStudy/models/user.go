package models;

import (
	"github.com/astaxie/beego/orm";

	// "github.com/astaxie/beego"
);

// Model Struct
// 如果有设置表前缀，那么表名就是前缀+表名；如果前缀设置为 prefix_ 那么表名为：prefix_auth_user
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


// auto
// 当 Field 类型为 int, int32, int64, uint, uint32, uint64 时，可以设置字段为自增健
// 当模型定义里没有主键时，符合上述类型且名称为 Id 的 Field 将被视为自增健。
// 鉴于 go 目前的设计，即使使用了 uint64，但你也不能存储到他的最大值。依然会作为 int64 处理。
// 数据库表默认为 NOT NULL，设置 null 代表 ALLOW NULL
// 还有更多orm关键字，如：index，unique，column，digits / decimals等等，详见文档，但更多的提倡是先直接通过sql进行建表，会更灵活些，且修改比较方便

/*PostgreSQL字段类型和go的数据类型对应关系
go	postgres
int, int32, int64, uint, uint32, uint64 - 设置 auto 或者名称为 Id 时	serial
bool	bool
string - 若没有指定 size 默认为 text	varchar(size)
string - 设置 type(char) 时	char(size)
string - 设置 type(text) 时	text
string - 设置 type(json) 时	json
string - 设置 type(jsonb) 时	jsonb
time.Time - 设置 type 为 date 时	date
time.Time	timestamp with time zone
byte	smallint CHECK(“column” >= 0 AND “column” <= 255)
rune	integer
int	integer
int8	smallint CHECK(“column” >= -127 AND “column” <= 128)
int16	smallint
int32	integer
int64	bigint
uint	bigint CHECK(“column” >= 0)
uint8	smallint CHECK(“column” >= 0 AND “column” <= 255)
uint16	integer CHECK(“column” >= 0)
uint32	bigint CHECK(“column” >= 0)
uint64	bigint CHECK(“column” >= 0)
float32	double precision
float64	double precision
float64 - 设置 digits, decimals 时	numeric(digits, decimals)*/
