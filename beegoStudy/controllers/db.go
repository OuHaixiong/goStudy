package controllers;

import (
	"fmt"
	"github.com/astaxie/beego";
	"beegoStudy/models";
	"github.com/astaxie/beego/orm";
	// "strings"
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
	var maps []orm.Params; // 下面是通过orm执行sql（原生查询）
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

// 下面测试表与表之间的关系，同时进行关系查询
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

	/*db, err := orm.GetDB();
    if (err != nil) {
        fmt.Println("get default DataBase error");
	} else {
		println(db); // db为一个对象，返回：0xc4200b2140
	}
	db, err = orm.GetDB("alias");
	if err != nil {
		fmt.Println("get alias DataBase error:", err); // get alias DataBase error: DataBase of alias name `alias` not found
	}*/

	// o2 := orm.NewOrm()
	// o2.Using("db2") // 切换数据库。默认使用 default 数据库，无需调用 Using

	// 下面演示关系查询
	/*o := orm.NewOrm();
	member := models.Member{Id:2}; // 注意这样的写法和new的写法是一样的，只不过在read传参时需要传引用，即:o.Read(&xx)
	err := o.Read(&member); // [SELECT "id", "name", "profile_id" FROM "member" WHERE "id" = $1 ] - `2`
	if (err == nil) {
		println(member.Name);
		if member.Profile != nil { // println(member.Profile) // 0xc42026a3e0
			// Member 和 Profile 是 OneToOne 的关系
			o.Read(member.Profile); // [SELECT "id", "age" FROM "profile" WHERE "id" = $1 ] - `2`
			println(member.Profile.Age); 
		}
	} else {
        println("不存在的记录");
	}*/
	// if (err == orm.ErrNoRows) { // 如果记录不存在，还可以这样写
	// 	println("不存在的记录");
	// }

	// 直接关联查询
	/*o := orm.NewOrm();
	member := models.Member{}; // 后面的{}一定不能少；相当于 var member models.Member;
	o.QueryTable("member").Filter("Id", 1).RelatedSel().One(&member); // Filter中查找的字段大小写均可。  .One(&xxx)参数一定只能按引用传递
	// 通过RelatedSel直接进行连表查询：
	// [SELECT T0."id", T0."name", T0."profile_id", T1."id", T1."age" FROM "member" T0 INNER JOIN "profile" T1 ON T1."id" = T0."profile_id" WHERE T0."id" = $1 LIMIT 1] - `1`
	println(member.Profile.Age); // 自动查询到 Profile
    println(member.Profile.Member.Name); // 因为在 Profile 里定义了反向关系的 User，所以 Profile 里的 User 也是自动赋值过的，可以直接取用。 */

	/*var profile models.Profile;
	o := orm.NewOrm();
	// Filter 字段筛选，大小写均可
	err := o.QueryTable("profile").Filter("Member__Id", 1).One(&profile); // 通过 Member 反向查询 Profile
	// [SELECT T0."id", T0."age" FROM "profile" T0 INNER JOIN "member" T1 ON T1."profile_id" = T0."id" WHERE T1."id" = $1 LIMIT 1] - `1`
	if (err == nil) {
		println(profile.Age)
		// fmt.Println(profile); // {1 30 <nil>}
	}*/
	
	/*o := orm.NewOrm();
	var articles []models.Article;
	// Article 和 Member 是 ManyToOne 关系，也就是 ForeignKey(外键) 为 Member
	num, err := o.QueryTable("article").Filter("Member", 3).RelatedSel().All(&articles);
	// [SELECT T0."id", T0."title", T0."member_id", T1."id", T1."name", T1."profile_id", T2."id", T2."age" FROM "article" T0 
	// INNER JOIN "member" T1 ON T1."id" = T0."member_id" INNER JOIN "profile" T2 ON T2."id" = T1."profile_id" WHERE T0."member_id" = $1 LIMIT 1000] - `3`
    if err == nil {
		fmt.Printf("%d member read \n", num);
        for _, article := range articles {
			fmt.Printf("Id: %d, Title: %s, Name:%s \n", article.Id, article.Title, article.Member.Name);
		}
	}*/

	// 同理 根据 Article.Title 也可以查询对应的 Member
	/*var member models.Member;
	o := orm.NewOrm();
	err := o.QueryTable("member").Filter("Article__Title", "欧妍妍发布的文章标题").Limit(1).One(&member);
	// [SELECT T0."id", T0."name", T0."profile_id" FROM "member" T0 INNER JOIN "article" T1 ON T1."member_id" = T0."id" WHERE T1."title" = $1 LIMIT 1] - `欧妍妍发布的文章标题`
	if err == nil {
		fmt.Println(member); // {3 欧妍妍 0xc420342b00 []}
	}*/

	// 下面演示多对多的关系（ManyToMany）【orm:"rel(m2m)"】
	/*o := orm.NewOrm();
	var articles []*models.Article;
	num, err := o.QueryTable("Article").Filter("tags__tag__name", "标签名").All(&articles); // .QueryTable()和.Filter() 的字段名可以大写也可以小写
	// [SELECT T0."id", T0."title", T0."member_id" FROM "article" T0 INNER JOIN "article_tags" T1 ON T1."article_id" = T0."id" 
	// INNER JOIN "tag" T2 ON T2."id" = T1."tag_id" WHERE T2."name" = $1 LIMIT 1000] - `标签名`
    if (err == nil) {
		println(num);
		for k, v := range articles {
			println(k, ": ", v.Id, " => ", v.Title) // 0 :  1  =>  欧妍妍发布的文章标题
		}
	} else { // 只有执行的sql出错了，才会到这来，不然不会到这来
		println("no records, num is:", num);
	}*/
	// 同理，也可以通过Article的title来查找tag
	/*o := orm.NewOrm();
	var tags []*models.Tag;
	num, err := o.QueryTable("tag").Filter("Articles__Article__Title", "欧妍妍发布的文章标题").All(&tags);
	// [SELECT T0."id", T0."name" FROM "tag" T0 INNER JOIN "article_tags" T1 ON T1."tag_id" = T0."id" 
	// INNER JOIN "article" T2 ON T2."id" = T1."article_id" WHERE T2."title" = $1 LIMIT 1000] - `欧妍妍发布的文章标题`
    if err == nil {
		println(num);
		for _, value := range tags {
			println(value.Name);
		}
	}*/

	// 下面是多对多（ManyToMany）关系载入
	/*o := orm.NewOrm();
	article := models.Article{Id:1};
	err := o.Read(&article); // [SELECT "id", "title", "member_id" FROM "article" WHERE "id" = $1 ] - `1`
	if (err == nil) {
		num, err := o.LoadRelated(&article, "Tags"); // 载入相应的 Tags
		// [SELECT T0."id", T0."name" FROM "tag" T0 INNER JOIN "article_tags" T1 ON T1."tag_id" = T0."id" WHERE T1."article_id" = $1 LIMIT 1000] - `1`
		if err == nil {
			println(num);
			for _, tag := range article.Tags {
				println(tag.Name);
			}
		}
	}*/
	// 同理也可以通过tag查找article
	o := orm.NewOrm();
	tag := models.Tag{Id:2};
	err := o.Read(&tag); // [SELECT "id", "name" FROM "tag" WHERE "id" = $1 ] - `2`
	if err == nil {
		num, err := o.LoadRelated(&tag, "Articles"); // LoadRelated 用于载入模型的关系字段，包括所有的 rel/reverse - one/many 关系
		// [SELECT T0."id", T0."title", T0."member_id" FROM "article" T0 INNER JOIN "article_tags" T1 ON T1."article_id" = T0."id" WHERE T1."tag_id" = $1 LIMIT 1000] - `2`
		if (err == nil) {
			println(num);
			for k, v := range tag.Articles {
				println(k, " => ", "id:", v.Id, ", title:", v.Title); // 1  =>  id: 2 , title: 文章标题2
			}
		}
	}

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

    /*users := []models.User{
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
	}*/

	// 下面演示多对多关系的插入
	/*o := orm.NewOrm();
	article := models.Article{Id:2};
	m2m := o.QueryM2M(&article, "Tags"); // 创建一个 QueryM2Mer 对象. 第一个参数的对象，主键必须有值, 第二个参数为对象需要操作的 M2M 字段
	tags := []*models.Tag{
		{Name:"名称1"},
		{Name:"名称2"},
	};
	for _, v := range tags {
		o.Insert(v); // [INSERT INTO "tag" ("name") VALUES ($1) RETURNING "id"] - `名称1`
		// o.Insert(&v);  // 不能这样写
	}
	num, err := m2m.Add(tags); // 多对多关系的添加，也可以写成：m2m.Add(tag1, tag2)。 Add 支持多种类型 Tag *Tag []*Tag []Tag []interface{}
	// [INSERT INTO "article_tags" ("article_id", "tag_id") VALUES ($1, $2), ($3, $4)] - `2`, `5`, `2`, `6`
	if err == nil {
		fmt.Println("Added numbs:", num);
	}*/

	o := orm.NewOrm();
	article := models.Article{Id:2};
	m2m := o.QueryM2M(&article, "Tags");
	nums, err := m2m.Count(); // 计算多对多数量： [SELECT COUNT(*) FROM "article_tags" T0 WHERE T0."article_id" = $1 ] - `2`
	if (err == nil) {
        fmt.Println("Total Numbs:", nums);
	}
	numbers, err := m2m.Clear(); // 清除所有多对多关系
	// [SELECT T0."id" FROM "article_tags" T0 WHERE T0."article_id" = $1 ] - `2`
	// [DELETE FROM "article_tags" WHERE "id" IN ($1, $2, $3)] - `3`, `4`, `5`
	if err == nil {
		fmt.Println("Removed Tag Numbers:", numbers);
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

	// 下面演示多对多关系删除
	/*tags := [] *models.Tag{
		{Id:5},
		{Id:6},
	};
	o := orm.NewOrm();
	article := models.Article{Id:2};
	m2m := o.QueryM2M(&article, "Tags");
	num, err := m2m.Remove(tags); 
	// [SELECT T0."id" FROM "article_tags" T0 WHERE T0."article_id" = $1 AND T0."tag_id" IN ($2, $3) ] - `2`, `5`, `6`
	// [DELETE FROM "article_tags" WHERE "id" IN ($1, $2)] - `6`, `7`
	if (err == nil) {
        fmt.Println("Removed nums:", num); // Removed nums: 2
	}*/

	c.Ctx.WriteString("测试db删除");
}

// 高级查询。测试 ORM 以 QuerySeter 来组织查询
func (c *DbController) Query_seter() {
	o := orm.NewOrm();
	// querySeter := o.QueryTable("member"); // orm.QueryTable(表名) ； 也可以直接使用对象做表名，如：u := new(models.User);o.QueryTable(u)
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

    // Exclude起到 排除条件 的作用，即相反，如：qs.Exclude("profile__id__exact", 8) 即 not profile.id=8 或 profile.id<>8

	/*var members []*models.Member;
	num, err := querySeter.OrderBy("-profile__age", "id").Limit(88).All(&members, "id", "Name", "profile"); // 需要查询的字段，大小写均可
	// SELECT T0."id", T0."name" FROM "member" T0 INNER JOIN "profile" T1 ON T1."id" = T0."profile_id" ORDER BY T1."age" DESC, T0."id" ASC LIMIT 88]
	// OrderBy() - 开头表示降序（DESC），反之没有表示升序（ASC）
	println("Number:", num, "; Error:", err); // Number: 3 ; Error: (0x0,0x0)
	for k, v := range members {
		println(k);
		println(v.Id, "=>", v.Name, "; Age:");
	}

	exists := o.QueryTable("user").Filter("name", "欧海雄007").Exist(); // [ SELECT COUNT(*) FROM "user" T0 WHERE T0."name" = $1 ] - `欧海雄007`
    fmt.Printf("Is exists: %s \n", exists); // Is exists: %!s(bool=true)*/

	// Distinct 去重（SELECT DISTINCT...）

	var maps []orm.Params;
	num, err := o.QueryTable("member").Values(&maps, "id", "name","profile","profile__age"); 
	// .Values() 返回结果集的 key => value 值，key 为 Model 里的 Field name，value 的值 以 string 保存
	if err == nil {
		fmt.Printf("Result Numbers: %d \n", num);
        for k, v := range maps {
			fmt.Println(k, "===>", v["Id"], v["Name"], v["Profile"], v["Profile__Age"]); // map 中的数据都是展开的，没有复杂的嵌套
			// 1 ===> 2 欧阳海雄 <nil> 35  特别注意：上面的字段必须大写。没有profile这个字段
		}
	}

	var lists []orm.ParamsList // .ValuesList() 和 .Values() 的区别是，.Values()的key是字段名，而.ValuesList()的key为数字
	number, err := o.QueryTable("member").ValuesList(&lists, "id", "name", "profile__age")
	if err == nil {
		fmt.Printf("Result Numbers: %d \n", number) //  Result Numbers: 3
		for _, row := range lists {
			// fmt.Printf("Name: %s， Age: %s\n", row[1], row[2]) // Name: 欧阳海雄， Age: %!s(int64=35)  类型不对时才返回这样
			fmt.Printf("Name: %s， Age: %d \n", row[1], row[2]) // Name: 欧阳海雄， Age: 35
		}
	}

	var list orm.ParamsList;
	n, err := o.QueryTable("user").ValuesFlat(&list, "name"); // 这里只允许传两个参数
	if err == nil {
		fmt.Printf("Result Numbs: %d \n", n);
		// fmt.Printf("All name and id is : %s \n", strings.Join(list, ", ")); // 这里有问题，不知道怎么才能显示
		// for _, l := range list {
			// println(list[0]);
		// }
	}

	// str := []string{"Hello", "World", "Good"}
    // fmt.Println(strings.Join(str, " ")) // 连接成字符串时，数组的每一项必须是字符串

    c.Ctx.WriteString("测试db高级查询");
}

// 原生sql执行
func (c *DbController) Protosomatic() {
	o := orm.NewOrm();
	var rawSeter orm.RawSeter; // 提倡 rawSeter := 写法
	rawSeter = o.Raw("update tag set name=? where id=?", "PHP", 3);  // [update tag set name=$1 where id=$2] - `PHP`, `3`
	// o.Raw("select name from member where id in(?,?,?)", ids)   // ids := []int{1,2,3}
	res, err := rawSeter.Exec(); // 运行sql语句
	// RawSeter.Exec() (sql.Result, error) 执行sql语句
	if err == nil {
		num, _ := res.RowsAffected();
		fmt.Println("sql exec row affected nums:", num); // sql exec row affected nums: 1
	}
	
	var tag models.Tag;
	e := o.Raw("SELECT id, name FROM tag WHERE id = ?", 3).QueryRow(&tag); // .QueryRow(interface{}) error : 查询一条记录，支持struct
	if e == nil {
		beego.Info(tag.Id, " => ", tag.Name);
	}
	// 如果查询所有的话，可以用.QueryRows；如：
	// type User struct {
	// 	Id       int
	// 	UserName string
	// }
	// var users []User
	// num, err := o.Raw("SELECT id, user_name FROM user WHERE id = ?", 1).QueryRows(&users)
	// if err == nil {
	// 	fmt.Println("user nums: ", num)
	// }
	

	var maps []orm.Params;
	r := o.Raw("select * from member where name=? and profile_id=?");
	result, err := r.SetArgs("Bear-Ou", 1).Values(&maps);
	if (err == nil && result > 0) {
		beego.Info("result:", result); // result: 1
		for key, item := range maps {
			beego.Info("key=>", key); // key=> 0
			beego.Info(item["id"], item["name"]); // 1 Bear-Ou 
		}
	}
	_, error := r.SetArgs("欧阳海雄", "2").Values(&maps); // .SetArgs(arg1,arg2,...) 可重复使用替换参数
	if error == nil {
		beego.Info(error); // <nil> 
		for _, item := range maps { // 在for中，第一个返回值如果舍弃掉，第二个是可以已定义过的，但如果在其他的地方，如直接调用函数返回值的话是不行的
			beego.Info(item["id"], item["name"]); // 2 欧阳海雄
		}
	}

	var list orm.ParamsList;
	num, err := o.Raw("select id from member where id < ?", 10).ValuesFlat(&list); // .ValuesFlat 返回单一字段数组值
	if err == nil && num > 0 {
		beego.Info(list); // [1 2 3 5]
        for _, id := range list {
			println(id); // (0x8e02e0,0xc42050c590)
			fmt.Println(id); // 1
		}
		// fmt.Println(strings.Join(list,",")) // 不能这样转换
	} 
    // 还有一些原生操作：Prepare、RowsToStruct、RowsToMap 详见：https://beego.me/docs/mvc/model/rawsql.md
    c.Ctx.WriteString("测试db的原生sql执行");
}

// QueryBuilder (构造查询)
func (c *DbController) Query_builder() {
    type User struct { // 结构体是可以写在函数里面的，如果写在里面，那么就只能在函数内使用了
		Name string
		Age int
	}

	var users []User
    // 我靠竟然只能用mysql，且用mysql生成sql竟然能在postgresql数据库中执行成功
	queryBuilder, err := orm.NewQueryBuilder("mysql") // 神奇了，我没有注册mysql的驱动，竟然能执行成功。(我猜应该是执行的sql语句是一样的，不然...)
	// queryBuilder, err := orm.NewQueryBuilder("postgres") // 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	if err != nil {
		beego.Error("无法通过驱动别名生成QueryBuilder对象");
		beego.Error(err.Error()); // unknown driver for query builder 
		// postgres query builder is not supported yet
		c.Ctx.WriteString("获取驱动出错了，无法created query builder；请查看日记");
		return
	}
	// 构建查询对象
	queryBuilder.Select("member.name", "profile.age").From("member").
		InnerJoin("profile").On("member.profile_id = profile.id").
		Where("profile.age > ?").
		OrderBy("name").Desc().
		Limit(10).Offset(0)
	sql := queryBuilder.String() // 导出 SQL 语句
	// [SELECT member.name, profile.age FROM member INNER JOIN profile ON member.profile_id = profile.id WHERE profile.age > $1 ORDER BY name DESC LIMIT 10 OFFSET 0] - `18`
	o := orm.NewOrm() // 执行 SQL 语句
	o.Raw(sql, 11).QueryRows(&users)
	for k, v := range users {
		beego.Info(k, "=> ", v.Name, v.Age) // 1 =>  Bear-Ou 30
	}
	

    c.Ctx.WriteString("测试db的QueryBuilder查询");
}

// 事务处理
func (c *DbController) Transaction() {
	/*o := orm.NewOrm()
	err := o.Begin()
	if (err != nil) {
		beego.Error("begin transaction is error:", err.Error())
		c.Ctx.WriteString("启动事务失败");
	}
	sql := "update profile set age=age+1"
	result, err := o.Raw(sql).Exec();// 运行sql语句
	if (err != nil) {
		fmt.Println(err.Error()) // pq: relation "profiles" does not exist
		err = o.Rollback()
		fmt.Println(err) // <nil>
		c.Ctx.WriteString("执行事务失败，已回滚");
	} else {
        number, _ := result.RowsAffected();
		err = o.Commit()
		fmt.Printf("sql exec row affected numbers: %d \n", number); // sql exec row affected numbers: 4
	}*/

	// 下面查看db的信息
	o := orm.NewOrm()
	// o.Using("databaseAlias")
	driver := o.Driver()
	fmt.Println(driver.Name()) // default
	driverType := driver.Type() // 4
	fmt.Println(driverType)
	fmt.Println(driverType == orm.DRPostgres) // true
	fmt.Println(driverType == orm.DRSqlite) // false

    c.Ctx.WriteString("测试db的事务处理");
}
