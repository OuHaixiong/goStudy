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
	/*o := orm.NewOrm(); // 创建Ormer对象
	o.Using("default"); // 默认使用default（可省略），你也可以指定为其他数据库
	profile := new(models.Profile);
	profile.Age = 35;

	member := new(models.Member);
	member.Profile = profile;
	member.Name = "欧阳海雄";

	fmt.Println(o.Insert(profile));
    fmt.Println(o.Insert(member));*/

	/*o := orm.NewOrm();
	member := new(models.Member);
	profile := new(models.Profile);
	tag := new(models.Tag);
	article := new(models.Article);
	profile.Age = 18;
	member.Profile = profile;
	member.Name = "欧妍妍";
	tag.Name = "标签名";
	article.Title = "欧妍妍发布的文章标题";
	article.Member = member;
	fmt.Println(o.Insert(profile));
	fmt.Println(o.Insert(member));
	fmt.Println(o.Insert(tag));
	// tags := []models.Tag{tag};
	// article.Tags = tags;
	fmt.Println(o.Insert(article));*/

	db, err := orm.GetDB();
    if (err != nil) {
        fmt.Println("get default DataBase error");
	} else {
		println(db); // db为一个对象，返回：0xc4200b2140
	}
	db, err = orm.GetDB("alias");
	if err != nil {
		fmt.Println("get alias DataBase error:", err); // get alias DataBase error: DataBase of alias name `alias` not found
	}

	// o2 := orm.NewOrm()
	// o2.Using("db2") // 切换数据库。默认使用 default 数据库，无需调用 Using

    c.Ctx.WriteString("OK!!!!");
}

func (c *DbController) Crud() {
	o := orm.NewOrm();
	user := new(models.User);
	user.Name = "ni好";
	user.Email = "nihao@xiqiyanyan.com";
	fmt.Println(o.Insert(user));  // 返回主键id和错误信息：5 <nil>
	user.Name = "wo ri";
	fmt.Println(o.Update(user));  // 返回更新条数和错误信息：1 <nil>
	fmt.Println(o.Read(user));    // 返回错误信息：<nil> 
	fmt.Println(o.Delete(user));  // 返回删除条数和错误信息：1 <nil>
	
	u := new (models.User);
	u.Id = 88;
	err := o.Read(u);
	if err != nil {
		fmt.Println("read error:", err); // read error: <QuerySeter> no row found
	}

	c.Ctx.WriteString("增删改查");
}

func (c *DbController) Read() {
	o := orm.NewOrm();
	user := models.User{Id:88};
	err := o.Read(&user); // Read 默认通过查询主键赋值
    if err == orm.ErrNoRows {
		fmt.Println("查询不到记录，Error：", err); // 查询不到记录，Error： <QuerySeter> no row found
	} else if err == orm.ErrMissPK {
        fmt.Println("找不到主键，Error：", err);
	} else {
		fmt.Println("已查找到记录，id：", user.Id, ", name:", user.Name, ", email:", user.Email);
	}

	u := models.User{Name:"ni好"};
	err = o.Read(&u, "name"); // 第二个参数是需要查询的字段，可大写也可小写
	// 相当于执行了SELECT "id", "name", "email" FROM "user" WHERE "name" = $1  但是会把查询到的第一条记录赋给对应的结构体对象
    if err == orm.ErrNoRows {
        fmt.Println("未找到记录");
	} else {
        fmt.Println(u.Id, "=>", u.Email);
	}

	var lists []orm.ParamsList;
	num, err := o.QueryTable("member").ValuesList(&lists, "id", "name", "profile__age");
    if err == nil {
		fmt.Printf("Result line numbers : %d \n", num);
		for _, row := range lists {
            fmt.Printf("Name:%s, Id:%s, Age:%s \n", row[1], row[0], row[2]);
		}
	}

	
	c.Ctx.WriteString("测试db读取");
}

func (c *DbController) Insert() {
	/*o := orm.NewOrm();
	u := models.User{Name:"Bear.Ou", Email:"258333309@163.com"};
	id, err := o.Insert(&u); // INSERT INTO "user" ("name", "email") VALUES ($1, $2) RETURNING "id"
	if err == nil {
		fmt.Println(id);
	} else {
        fmt.Println("插入数据失败，Error:", err);
	}*/

    users := []models.User{
		{Name:"Bear01", Email:"xx1@bb.com"},
		{Name:"Bear02"},
		{Name:"Bear03", Email:"xx3@bb.com"},  // 这里的逗号不能少，少了默认就是分号（;）了，语法就错了
	};
	o := orm.NewOrm();
	successNumbers, err := o.InsertMulti(100, users); // func (o *orm) InsertMulti(bulk int, mds interface{}) (int64, error)
	// InsertMulti 第一个参数为并列插入的数量，第二个参数为结构体对象的切片（slice）; 返回成功插入的条数和错误信息
	// [INSERT INTO "user" ("name", "email") VALUES ($1, $2), ($3, $4), ($5, $6)] - `Bear01`, `xx1@bb.com`, `Bear02`, ``, `Bear03`, `xx3@bb.com`
	if (err != nil) {
		fmt.Println("插入多条数据失败，Error：", err);
	} else {
        fmt.Println("成功插入 ", successNumbers, " 条数据");
	}

	

    c.Ctx.WriteString("测试db插入");
}

func (c *DbController) Update() {
	o := orm.NewOrm();
	user := models.User{Id:11};
    if o.Read(&user) == nil {
		user.Name = "欧海雄007";
		if num, err := o.Update(&user); err == nil {
			// UPDATE "user" SET "name" = $1, "email" = $2 WHERE "id" = $3] - `欧海雄007`, `258333309@163.com`, `11`
			// o.Update(&user, "field1", "Field2", ...); Update默认更新所有的字段，可以只更新指定的字段
            fmt.Println(num); // Update 成功返回更新的条数和错误信息
		}
	}

	// 下面演示根据条件，批量更新记录
	num, err := o.QueryTable("user").Filter("name__startswith", "欧").Update(orm.Params{"name":"Bear--海海"});
	// [UPDATE "user" SET "name" = $1 WHERE "id" IN ( SELECT T0."id" FROM "user" T0 WHERE T0."name"::text LIKE $2  )] - `Bear--海海`, `欧%`
	fmt.Printf("Affected line num: %s, error: %s \n", num, err); // Affected line num: %!s(int64=3), error: %!s(<nil>)

	// num, err := o.QueryTable("user").Update(orm.Params{
	// 	"nums": orm.ColValue(orm.ColAdd, 100), // SET nums = nums + 100 原子操作，增加字段值 orm.ColValue 支持以下操作：ColAdd：加，ColMinus：减，ColMultiply：乘，ColExcept：除
	// })

	c.Ctx.WriteString("测试db更新");
}

func (c *DbController) Delete() {
	o := orm.NewOrm();
	u := models.User{Id:15}
    if num, err := o.Delete(&u); err == nil {
        fmt.Println("删除成功，删除了 ", num, " 条数据");
	} else { // 如果记录不存在，sql也是执行成功的，err也是为nil。除非sql执行失败才会返回err
		fmt.Println("删除失败");
	}
	// SELECT T0."id" FROM "post" T0 WHERE T0."user_id" IN ($1) 
	// [DELETE FROM "post" WHERE "id" IN ($1, $2)] - `6`, `7`
	// Delete 操作会对反向关系进行操作，此例中 Post 拥有一个到 User 的外键。删除 User 的时候。如果 on_delete 设置为默认的级联操作，将删除对应的 Post
	// 即如果设置了 User   *User   `orm:"rel(fk)"`; （外键）就会一并删除

	// 下面演示批量删除功能
	num, err := o.QueryTable("user").Filter("name__contains", "ni").Delete();
	// [SELECT T0."id" FROM "user" T0 WHERE T0."name"::text LIKE $1 ] - `%ni%`
    // [DELETE FROM "user" WHERE "id" IN ($1)] - `10`
    // [SELECT T0."id" FROM "post" T0 WHERE T0."user_id" IN ($1) ] - `10`
	fmt.Printf("Affected line num: %s, Error: %s \n", num, err); // Affected line num: %!s(int64=1), Error: %!s(<nil>) 



	c.Ctx.WriteString("测试db删除");
}

// 高级查询。测试 ORM 以 QuerySeter 来组织查询
func (c *DbController) Query_seter() {
	o := orm.NewOrm();
	querySeter := o.QueryTable("member"); // orm.QueryTable(表名) ； 也可以直接使用对象做表名，如：u := new(models.User);o.QueryTable(u)
	// var m models.Member;

	// err := querySeter.Filter("name", "熊熊").One(&m, "Name", "Id"); // 条件查询:SELECT T0."name" FROM "member" T0 WHERE T0."name" = $1 LIMIT 1
	// if err == orm.ErrMultiRows { // 如果有多条记录，这里并不会报错
    //    fmt.Printf("Returned multi rows no one"); // 找到多条记录的时候报错
	// } else if err == orm.ErrNoRows {
    //     fmt.Printf("Not row found"); // 没有找到记录
	// }
	// // println(m.Id, " ", m.Name, " ", m.Profile.Age);
	// println(m.Id, " ", m.Name);

    // err := querySeter.Filter("profile__age", 18).One(&m); // SELECT T0."id", T0."name", T0."profile_id" FROM "member" T0 INNER JOIN "profile" T1 ON T1."id" = T0."profile_id" WHERE T1."age" = $1 LIMIT 1
	// qs.Filter("profile__age__gt", 18) // WHERE profile.age > 18
    // qs.Filter("profile__age__gte", 18) // WHERE profile.age >= 18
	// qs.Filter("profile__age__in", 18, 20) // WHERE profile.age IN (18, 20) 或 ids:=[]int{17,18,19,20}; qs.Filter("profile__age__in", ids)
	// 操作符：
	// exact / iexact 等于
	// contains / icontains 包含 WHERE name LIKE BINARY '%slene%' /  WHERE name LIKE '%slene%'
	// gt / gte 大于 / 大于等于
	// lt / lte 小于 / 小于等于
	// startswith / istartswith 以…起始； qs.Filter("name__startswith", "slene") => WHERE name LIKE BINARY 'slene%'
	// endswith / iendswith 以…结束；     qs.Filter("name__iendswithi", "slene") => WHERE name LIKE '%slene'
	// in
	// isnull ==  qs.Filter("profile_id", nil) // WHERE profile_id IS NULL    qs.Filter("profile_id__isnull", true) => WHERE profile_id IS NULL 如果后面的参数为false则为is not null
	// 后面以 i 开头的表示：大小写不敏感
	// if err == orm.ErrNoRows {
	//     fmt.Printf("Not row found"); // 没有找到记录
	// }
	// println(m.Id, " ", m.Name, " ", m.Profile.Age);

	/*querySeter = querySeter.Filter("name__istartswith", "欧").Filter("profile__age__gte", 18); // where name like '欧%' and profile.age >= 18
	// [ORM]2018/09/28 05:51:24  -[Queries/default] - [  OK /    db.Query /     1.1ms] - [SELECT T0."id", T0."name", T0."profile_id" FROM "member" T0 INNER JOIN "profile" T1 ON T1."id" = T0."profile_id" WHERE UPPER(T0."name"::text) LIKE UPPER($1) AND T1."age" >= $2 LIMIT 1000] - `欧%`, `18`
	var members []*models.Member;
	num, err := querySeter.All(&members); // All / Values / ValuesList / ValuesFlat 受到 Limit 的限制，默认最大行数为 1000
	fmt.Printf("Returned rows Num: %s, Error:%s\n", num, err); // Returned rows Num: %!s(int64=2), Error:%!s(<nil>)
    for _, value := range members {
		println(value.Id, "=>", value.Name);
		// 2 => 欧阳海雄
		// 3 => 欧妍妍
	}*/

	/*condition := orm.NewCondition();
	ids := []int{88,89,90,1,2,3,5}; 
	cond1 := condition.And("name__isnull", false).AndNot("id__in", ids).Or("profile__age__lt", 100); // 条件not...表示相反，如not id=1即id!=1, not id in(1,2) 即 id not in(1,2)[id不等于1,2的]
    // [SELECT COUNT(*) FROM "member" T0 INNER JOIN "profile" T1 ON T1."id" = T0."profile_id" WHERE T0."name" IS NOT NULL AND NOT T0."id" IN ($1, $2, $3) OR T1."age" < $4 ] - `88`, `89`, `90`, `100`
	qs := o.QueryTable("member");
	qs = qs.SetCond(cond1);
	number, _ := qs.Count();
    println("共 ", number, " 条记录");
	// cond2 := condition.AndCond(cond1).OrCond(condition.And("name", "海海"));  // where (...) or (name="海海")
	// num, err := qs.SetCond(cond2).Count();*/

	// var members []*models.Member;
	// num, err := querySeter.Limit(1, 2).All(&members, "Id", "name"); // SELECT T0."id", T0."name" FROM "member" T0 LIMIT 1 OFFSET 2
	// QuerySeter.Limit(limit, offset) 默认select查询的最大行数为1000；如果为Limit(-1)即为没有limit（no limit）。也可以单独设置offset：.Offset(20) LIMIT 1000 OFFSET 20
	// 分组用QuerySeter.GroupBy("id", "age"); GROUP BY id,age

	var members []*models.Member;
	num, err := querySeter.OrderBy("-profile__age", "id").Limit(88).All(&members, "id", "Name", "profile"); // 需要查询的字段，大小写均可
	// SELECT T0."id", T0."name" FROM "member" T0 INNER JOIN "profile" T1 ON T1."id" = T0."profile_id" ORDER BY T1."age" DESC, T0."id" ASC LIMIT 88]
	println("Number:", num, "; Error:", err); // Number: 3 ; Error: (0x0,0x0)
	for k, v := range members {
		println(k);
		println(v.Id, "=>", v.Name, "; Age:");
	}

	exists := o.QueryTable("user").Filter("name", "欧海雄007").Exist(); // [ SELECT COUNT(*) FROM "user" T0 WHERE T0."name" = $1 ] - `欧海雄007`
    fmt.Printf("Is exists: %s \n", exists); // Is exists: %!s(bool=true)


    c.Ctx.WriteString("测试db高级查询");
}
