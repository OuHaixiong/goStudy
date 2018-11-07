// 练习使用gorm 类似的还有xorm：https://github.com/go-xorm/xorm/blob/master/README_CN.md
// 还有增强版的xorm：https://github.com/xormplus/xorm       
// 还有beego的orm
// 如果仅仅使用mysql的话，还是didi的gendry:https://github.com/didi/gendry/blob/master/translation/zhcn/README.md

// gorm中文文档：https://github.com/jasperxu/gorm-cn-doc (http://gorm.book.jasperxu.com/)

package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //引入驱动
	"time"
)

var db *gorm.DB // 已经进行了db的初始化操作,db为全局变量

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=root dbname=testdb sslmode=disable password=e12345ooo6"); // 无法连接数据库时，程序一直处于卡死状态
	// 这是什么鬼，有bug吧，为什么密码不对也能连上(太搞笑了，坑爹呀)
	if (err != nil) {
        println("Error:", err.Error()) // Error: dial tcp 127.0.0.1:54321: connect: connection refused
	}
	defer db.Close() // 这里的db为局部变量和全局变量的db并不冲突

	db.LogMode(true) // 启用Logger，显示详细日志

	boolean := db.HasTable(&User{}); // 检查模型`User`表是否存在
	println(boolean)
	boolean = db.HasTable("user") // 检查表`user`是否存在
	println(boolean)

    // 创建表
	// db.CreateTable(&UserInfo{}) // 如果表已经存在，不会再创建
	// CREATE TABLE "user_infos" ("id" bigserial,"name" varchar(100) DEFAULT 'ououou',
	// "age" integer,"birthday" timestamp with time zone,"created_at" timestamp with time zone,
	// "updated_at" timestamp with time zone , PRIMARY KEY ("id"))

    // var userInfo = UserInfo{
	// 	Name : "欧阳海雄2",
	// 	Age  : 33,
	// }
	// 插入数据
	// db.Create(&userInfo) // INSERT INTO "user_infos" ("name","age","birthday","created_at","updated_at") VALUES 
	    // ('欧阳海雄','35','0001-01-01 00:00:00','2018-11-07 02:11:28','2018-11-07 02:11:28') RETURNING "user_infos"."id"
	// var birthday = time.Date(1983, time.Month(8), 15, 14, 20, 2, 0, time.Now().Location())
	// 根据指定时间返回time.Time；参数分别为：年，月，日，时，分，秒，纳秒，时区
    // var userInfo2 = UserInfo {
	// 	Name      : "ou海海",
	// 	Age       : 28,
	// 	Birthday  : birthday,
	// }
	// db.Create(&userInfo2) // INSERT INTO "user_infos" ("name","age","birthday","created_at","updated_at") VALUES 
	    // ('ou海海','28','1983-08-15 14:20:02','2018-11-07 02:30:29','2018-11-07 02:30:29') RETURNING "user_infos"."id"

	// 下面演示读取数据
	/* var userInfo UserInfo
	db.First(&userInfo) // 获取第一条记录，按主键正序排序：SELECT * FROM "user_infos"   ORDER BY "user_infos"."id" ASC LIMIT 1
	println(userInfo.ID, userInfo.Name);
	// 特别注意了，不能进行连续查询，因为这个对象已经赋值了
	var lastUserInfo UserInfo
	db.Last(&lastUserInfo) // 获取最后一条记录，按主键倒序排序：SELECT * FROM "user_infos"   ORDER BY "user_infos"."id" DESC LIMIT 1
	println(lastUserInfo.ID, lastUserInfo.Age)
	var users []UserInfo
	db = db.Find(&users) // 获取所有记录: SELECT * FROM "user_infos"
	// if err != nil { // 返回的是db数据库对象，并非是错误信息
	// 	println(err)
	// }
	for k, v := range users {
		println(k)
		println(v.Name, v.Age)
	}
	var u2 UserInfo
	db.First(&u2, 3) // 使用主键获取记录：SELECT * FROM "user_infos"  WHERE ("user_infos"."id" = 3) ORDER BY "user_infos"."id" ASC LIMIT 1
	println(u2.ID, u2.Name, u2.Age, u2.IgnoreMe, u2.Birthday.String()) // 如果查询记录不存在的话，对象返回的均为默认值 */

	// 下面演示删除记录
	var u4 UserInfo
	db.Table("user_infos").Select("id, name, age").Where("id = ?", 2).Scan(&u4) // SELECT id, name, age FROM "user_infos"  WHERE (id = '2') 
	println(u4.ID, u4.Name, u4.Age)
	if (u4.ID > 0) {
		db = db.Delete(&u4) // DELETE FROM "user_infos"  WHERE "user_infos"."id" = '2'  [1 rows affected or returned ]
		println(db)
		// 特别注意： 删除记录时，需要确保其主要字段具有值，GORM将使用主键删除记录，如果主要字段为空，GORM将删除模型的所有记录
		println("删除记录条数：", db.RowsAffected)
	}

	// 下面演示修改记录
	// var u5 UserInfo
	// db.Where("name=?", "ou海海").First(&u5) // 获取第一个匹配记录
	// // SELECT * FROM "user_infos"  WHERE (name='ou海海') ORDER BY "user_infos"."id" ASC LIMIT 1
	// db.Model(&u5).Update("name", "欧海雄") // UPDATE "user_infos" SET "name" = '欧海雄', "updated_at" = '2018-11-07 03:50:02'  WHERE "user_infos"."id" = '3'
	// // 如果表中有updated_at这个字段（不管结构体中的字段名是否为UpdatedAt），更新时会自动更新该字段的时间
	// println(db.RowsAffected) // 奇怪的是更新了一条记录，这里返回的是0
	// println(u5.ID, u5.Name) // 3 欧海雄
	// db.Model(&u5).Update("created_at", time.Now()) // UPDATE "user_infos" SET "created_at" = '2018-11-07 03:56:54', "updated_at" = '2018-11-07 03:56:54'
	// 我靠，特别注意上面的执行呀，上面这样执行更新的是所有记录（坑还挺多）， 我我靠，如果查询的记录不存在的话，更新的也是全部的数据，你妈的，坑呀
	// 还有，更新字段用的是表字段名可以，用结构体的属性名也行

	// 下面演示原生语句的执行
	var u6 UserInfo
	db.Raw("select name, age from user_infos where age >= ?", 18).Scan(&u6) // Raw SQL 原生sql执行
	println(u6.ID, u6.Name, u6.Age)
	var userList []UserInfo
	db.Raw("select id, name, age from user_infos where age >= ?", 18).Scan(&userList)
	for _, v := range userList {
		println(v.ID, v.Name, v.Age)
	}
	// 注意上面的两种写法，如果传单个结构，查找的结果的最后一个会赋值给它，如果传的是数组（数组的结构体），那么所有的记录将赋值给数组
	
	var result Result
    db.Raw("select sum(age) as total_age from user_infos where age>=18").Scan(&result);
	println(result.TotalAge)
}

type User struct { // 特别注意了：在gorm中，表名和结构体的名称是一样的，如果需要单独返回表名，需要实现TableName方法
	Id uint `gorm:"primary_key"`       // int64 
	Name string
	Email string
}

func (User) TableName() string { // 返回表名
	return "user"
}

type UserInfo struct { // 默认表名是结构体名称的复数形式，即后面加了个s，除第一个字母大写外，遇到大写字母就转小写并在前面加“_”； 如：user_infos
	ID int64 // 字段`ID` 默认为主键；在表的字段名为：id。也可以使用 `gorm:primary_key` 来设置主键
	// 默认主键是自增的，所有可以不需要设置； 如果需要设置自增，执行 ： `gorm:"AUTO_INCREMENT"` 。对应postgresql：bigserial, PRIMARY KEY ("ID")
	Name string `gorm:"default:'ououou';size:100"`  // string 默认长度255。对应postgresql：varahar(100) default 'ououou'
	Age uint // 对应postgresql：integer
	Birthday time.Time // 对应postgresql：timestamp with time zone
	CreatedAt time.Time // 列名是字段名的蛇形小写；即：`created_at`。 字段CreatedAt用于存储记录的创建时间，将被设置为当前时间
	// 对应postgresql："created_at" timestamp with time zone
	UpdatedTime time.Time `gorm:"column:updated_at"` // 重设列名。 对应postgresql："updated_at" timestamp with time zone, 
	// 保存具有UpdatedAt字段的记录将被设置为当前时间,用于存储记录的修改时间
	IgnoreMe int `gorm:"-"` // - 表示忽略这个字段

}

type Result struct {
	TotalAge int
}
