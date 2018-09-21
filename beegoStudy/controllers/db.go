package controllers;

import (
	"fmt"
	"github.com/astaxie/beego";
	"beegoStudy/models";
    "github.com/astaxie/beego/orm";
);

type DbController struct {
	beego.Controller;
}

// func init() { // 如果需要初始化一个东东的时候，就写在这里
// }

// 测试数据库的操作
func (c *DbController) Test() {
	/*o := orm.NewOrm();
	user := models.User{Name:"欧海雄-Bear", Email:"258333309@qq.com"};
	id, err := o.Insert(&user); // insert 插入一条记录
	beego.Info(id);
	fmt.Printf("ID:%d, ERR:%v \n", id, err); // ID:2, ERR:<nil>

	user.Name = "ououou";
	num, err := o.Update(&user); // update 更新一条数据
    fmt.Printf("NUM: %d, ERR: %v\n", num, err);*/

	// read one 读取一条记录
	/*o := orm.NewOrm();
	tb1 := models.Tb1{Id:101}; // 声明一个tb1的结构体，并赋初值
	err := o.Read(&tb1);
	fmt.Printf("read is ERR: %v \n", err); // 读取成功返回：read is ERR: <nil>
    println(tb1.Name);*/

	/*o := orm.NewOrm();
	u := models.User{Id:3};
	num, err := o.Delete(&u); // delete 删除一条记录
	fmt.Printf("NUM:%d, ERR:%v\n", num, err);*/

	o := orm.NewOrm();
	var posts []*models.Post; // declare slice 声明切片，类型为Post结构体
	qs := o.QueryTable("post"); // 要查询的表
	num, err := qs.Filter("User__Name", "欧阳海雄").All(&posts); // 连表查询,查询条件是： 表user的字段name中值为欧阳海雄的记录。查询所有并赋值给切片posts
	// SELECT T0."id", T0."title", T0."user_id" FROM "post" T0 INNER JOIN "user" T1 ON T1."id" = T0."user_id" WHERE T1."name" = $1 LIMIT 1000] - `欧阳海雄`
	beego.Info(num, err);
	beego.Info(posts); // 2018/09/20 23:46:27.750 [I] [db.go:47] [0xc42044e360 0xc42044e3a0]  如果为两条记录的话，切片中存储的是连个内存地址
	for k, v := range posts {
		beego.Info(k, "=>Id:", v.Id);
		beego.Info(k, "=>Title:", v.Title);
		beego.Info(k, "=>User:", v.User); // 0 =>User: &{1  }  为结构体
		beego.Info(k, "=>user_id:", v.User.Id); // 2018/09/21 00:00:21.616 [I] [db.go:52] 1 =>user_id: 1 注意这里的写法 
	}

    o = orm.NewOrm();
	var maps []orm.Params; // 下面是通过orm执行sql
	num, err = o.Raw(`select u.*,p.title from "user" u, "post" p where p."user_id"=u."id"`).Values(&maps); // inner join  连表查询，注意表名最好用""括起来
    for key, term := range maps {
		fmt.Printf("%d => %v", key, term); // 返回map结构的数据：1 => map[id:1 name:欧阳海雄 email:Bear@maimengmei.com title:hehe]
		beego.Info(term["name"], term["title"]);
	}

	// 下面演示，事务处理
	/*o.Begin();
	u := models.User{Id:2};
	p := models.Post{Title:"这个是标题", User:&u};
	id, err := o.Insert(&p);
	if (err == nil) {
		beego.Info(id);
		o.Commit();
	} else {
        o.Rollback();
	}*/

    c.Ctx.WriteString("It is OK!"); // 输出字符串
}

// 下面测试表与表之间的关系
func (c *DbController) Relation() {
	o := orm.NewOrm();
    o.Using("default"); // 默认使用default（可省略），你也可以指定为其他数据库


	o := orm.NewOrm()
    o.Using("default") // 默认使用 default，你可以指定为其他数据库

    profile := new(Profile)
    profile.Age = 30

    user := new(User)
    user.Profile = profile
    user.Name = "slene"

    fmt.Println(o.Insert(profile))
    fmt.Println(o.Insert(user))
}