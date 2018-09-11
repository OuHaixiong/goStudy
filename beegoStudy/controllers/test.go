package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"fmt"
	"strconv"

)

type TestController struct {
	beego.Controller
}

func (c *TestController) Index() {
	c.Ctx.WriteString("这里是test控制器下面的Index方法 ");
	// c.Ctx.WriteString(c.Ctx.Input.Param(":controller")); // 没有:controller和:method这两个变量
	// c.Ctx.WriteString(c.Ctx.Input.Param(":method"));
	params := c.Ctx.Input.Params(); // GET /test/index/2048/09/05/bear_007.html 
	println(params) // params为0xc42029b2c0 : 这个为map集合数据的首存储地址
    for k, v := range params{ // 除了前缀两个 /:controller/:method 的匹配之外，剩下的 url beego 会帮你自动化解析为参数，保存在 this.Ctx.Input.Params 当中
		println(k, "=>", v)
		/* :splat => 2048/09/05/bear_007.html
           0 => 2048
           1 => 09
           2 => 05
           3 => bear_007.html */
	}
}

func (c *TestController) Api() {
	id := c.Ctx.Input.Param(":id");
	println(id)
	c.Ctx.WriteString(id);
}

func (c *TestController) Prepare() { // 执行Method方法之前会执行Prepare方法，用户可以重写这个函数实现类似用户验证之类
	println("====每个method前都会先调用Prepare方法====");
	// 假如用户认证不通过，输出信息并接受执行逻辑
	// c.Data["json"] = map[string]interface{}{"name":"ouyanhaixiong"}
	// c.ServeJSON() // 输出为json格式的字符串
	// c.StopRun();
}

func (c *TestController) Finish() { // 执行完Method方法后或执行Finish方法，用户可以执行例如数据库关闭，清理数据之类的工作。
    println("====每个method执行后都会调用Finish方法====");
}

func (c *TestController) Download() {
	// path := c.Ctx.Input.Param(":path"); // 获取路由定义的变量
	// ext := c.Ctx.Input.Param(":ext");
	// c.Ctx.WriteString(":path为" + path); // 获取路径；如file/api
	// c.Ctx.WriteString(":ext为" + ext); // 获取后缀名；如html
	splat := c.Ctx.Input.Param(":splat");
	c.Ctx.WriteString(":splat为" + splat);
}

// @router /test/abc/:key [get]
func (c *TestController) Bb() { // 上面的注释是有用的，用来注册路由。效果和自己通过 Router 函数注册是一样的
	c.Ctx.WriteString("注解注册的路由");
}

// 测试form表单，视图文件渲染
func (c *TestController) FormRend() { // 貌似go语言方法不区分大小写，所以如果多个单词的方法，建议用中下划线隔开(不支持中划线)
	v := c.GetSession("refreshNumber"); // 获取session
    if v == nil {
		c.SetSession("refreshNumber", int(1)); // 设置session
		c.Data["num"] = 0;
	} else {
		c.SetSession("refreshNumber", v.(int) + 1);
		c.Data["num"] = v.(int)
	}

	// println("Session ID:", sess.SessionID()) // TODO 无法获取

    c.TplName = "testController/formRend.html"  // 模板文件默认为：testcontroller/formrend.tpl
}

func (c *TestController) FormPost() { // 测试form表单post提交数据
	message := c.GetString("message"); // 获取post提交的字符串
	// id64, err := c.GetInt("id"); // GetInt(key string) (int64, error)    这样写是正确的，对前台传过来的数字直接转为int64了
	// 特别注意了：c.GetInt(key String) 其实也是调用strconv.Atoi方法，如果为空字符串，会直接跑挂程序
	// if err != nil {
    //     log.Fatal(err)
	// }
	// fmt.Printf("id 的类型为： %T \n", id64); // id 的类型为： int   打印一个变量的类型

	id := c.Input().Get("id"); // .Input().Get 获取到的参数为字符型
	fmt.Printf("id 的类型为： %T \n", id); // id 的类型为： string 
	var idString string;
    if (id == "") { // 如果id没有传值，为空字符串
        idString = "0";
	} else {
        idString = id;
	}

    // go 没有 try{}catch
	idInt64, err := strconv.Atoi(idString); // strconv.Atoi(String xxx) 字符串转整型
	// 特别注意了，上面的这行程序，如果转换出错了，程序进程直接就跑死了
	if err != nil { // 出错跑死了，这行代码根本没有用
        println("转换为整型出错了：", err);
	}
	fmt.Printf("id 的类型为： %T \n", idInt64); // id 的类型为： int

	boolean, _ := c.GetBool("boolean"); // GetBool(key string) (bool, error) 获取参数的boolean
	println("boolean为：", boolean); // 可以把字符串的true和false转为布尔值
	fmt.Printf("boolean 的类型为： %T \n", boolean); // boolean 的类型为： bool

	c.Ctx.WriteString(message);
	if (boolean) {
		c.Ctx.WriteString(" 选择的是正确：true");
	} else {
		c.Ctx.WriteString(" 选择的是错误：false");
	}
	// c.Ctx.WriteString(id64); // 这里无法编译通过（返回的是int64位数字，不能转为字符串）：cannot use id64 (type int) as type string in argument to c.Controller.Ctx.WriteString
	
	// 下面演示通过结构体来获取参数
	userStruct := user{}
	if err := c.ParseForm(&userStruct); err != nil {
		log.Fatal(err)
	}
	println(userStruct.Id);
	println(userStruct.Name); // 如果传入欧阳海雄：(0x840a20,0xc420332140)
	if name1, ok1 := userStruct.Name.(string); ok1 { // 对应任意类型的变量，需要先判断类型后才能打印（使用）
		println("任意类型的参数Name值为：", name1)
	} else if name2, ok2 := userStruct.Name.(int); ok2 {
        println(name2)
	}
	println(userStruct.Age);
	println(userStruct.Email);

	
}

type user struct {
	Id      int            `form:"-"`  // 或者为空不写也是可以的（但需要表单中无此字段名）
	Name    interface{}    `form:"username"`  // interface{}:代表为任意类型的数据
	Age     int            `form:"age"` // 如果表单传入的不是数字，程序将跑挂。这里的字符串转整用的是strconv.ParseInt
	// Email   string         `form:"email"` // 表单字段大小写是有区别的
	Email   string  // 如果没有写，就会把表单字段和结构体属性字段同名的对应起来
}


func (c *TestController) Form_body() { // 函数名不能用中划线（-）
	// var ob models.Object  // 有问题？？？
	// var err error
	// if err = json.Unmarshal(c.Ctx.Input.RequestBody, &ob); err == nil {
	// 	objectId := models.AddOne(ob);
	// 	c.Data["json"] = "{\"ObjectId\":\"" + objectId + "\"}"
	// } else {
    //     c.Data["json"] = err.Error()
	// }

	body := c.Ctx.Input.RequestBody;
	fmt.Printf("body 的类型为： %T \n", body); // []uint8 ： 二进制流数据
	for k, v := range body { // 0 => 117
		// 1 => 115 ...
		println(k, "=>", v); // 看起来像二进制数据呀
	}
	println(string(body)); // 二进制流数组转为字符串。返回：username=%E6%AC%A7%E9%98%B3%E6%B5%B7%E9%9B%84&age=35&Email=Bear%40maimengmei.com  即 欧阳海雄、35、Bear@maimengmei.com

	// c.Data["json"] = map[string]interface{}{"name":"欧海雄"}; //格式化为json字符串
	// 第二种直接声明：
	type my struct {
		name string
		age int
	}
	myStruct := my{"欧阳海雄", 35};
	println(myStruct.name);
	println(myStruct.age);
	c.Data["json"] = &myStruct // 这样写无法输出json
	c.ServeJSON(); // 输出为json的字符串。 调用 ServeJSON 之后，会设置 content-type 为 application/json，然后同时把数据进行 JSON 序列化输出。
    c.ServeJSONP(); // 调用 ServeJSONP 之后，会设置 content-type 为 application/javascript，然后同时把数据进行 JSON 序列化，然后根据请求的 callback 参数设置 jsonp 输出。
}

func (c *TestController) Form_upload() { // 测试上传文件
	// f, h, err := c.GetFile("upload_name") //获取上传文件信息
	// Controller.GetFile(key string) (multipart.File, *multipart.FileHeader, error)  读取上传的文件信息，用户可以根据这些信息来处理上传的文件，如：过滤、保存等
	// defer f.Close() // 延迟关闭打开的文件句柄
	// if err != nil {
	// 	log.Fatal("use GetFile func by upload file is err:", err);
	// }
	//c.SaveToFile("upload_name", "static/upload/" + h.Filename); // 保存位置在 static/upload, 没有文件夹要先创建
    c.SaveToFile("upload_name", "/data/www/go/src/beegoStudy/static/upload/abc.jpeg"); // 第一个参数为form表单的文件上传的name，第二个参数为保存的完整路径，可为相对路径也可绝对路径
	// Controller.SaveToFile(form_file_name string, to_file_path string) error 实现快速保存文件的功能
	c.Ctx.WriteString("上传文件成功！");
}

// 测试数据绑定
func (c *TestController) Bind() { // /test/bind.html?id=123&isok=true&ft=1.2&ol[0]=1&ol[1]=2&ul[]=str&ul[]=array&user.Name=astaxie
	var id int;
	c.Ctx.Input.Bind(&id, "id"); // 对url中请求的参数进行数据绑定
	println(id); // id=123

	var boolean bool;
	c.Ctx.Input.Bind(&boolean, "isok");
	println(boolean) // boolean=true

	var f64 float64;
	c.Ctx.Input.Bind(&f64, "ft"); // f64=1.2
	println(f64); // 打印为：+1.200000e+000

	ol := make([]int, 0, 2)
	c.Ctx.Input.Bind(&ol, "ol");
	// println(ol) // 打印为：[2/2]0xc420024140  因为是切片
	for k, v := range ol {
		println(k, " => ", v); // 0  =>  1   ,    1  =>  2
	}

	ul := make([]string, 0, 2);
	c.Ctx.Input.Bind(&ul, "ul");
	println(ul); // 打印为：[2/2]0xc42000c220
	for k, v := range ul {
		println(k, " => ", v); // 0  =>  str ,   1  =>  array
	}

	var u user
	c.Ctx.Input.Bind(&u, "user");
	// println(u.Name); // Name的类型（type）为： interface {} 任何参数类型 ；打印出来为：(0x0,0x0)
	// println(string(u.Name)); // 这样写是错误的：cannot convert u.Name (type interface {}) to type string: need type assertion
    if name1, ok1 := u.Name.(string); ok1 { // 真晕，没有赋值成功，不知道咋弄
		println(name1)
	} else if name2, ok2 := u.Name.(int); ok2 {
        println(name2)
	}

	c.Ctx.WriteString("绑定完成！");
}

// beego.InsertFilter("/*", beego.BeforeRouter, FilterUserLogin); // 没编译成功；原因是不能在函数外，调用有返回值的函数

func (c *TestController) Login() {
	c.TplName = "testController/login.html"  // 模板文件默认为：testcontroller/formrend.tpl
}
