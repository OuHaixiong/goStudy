package main

import (
    // abcP "./abc" // 两种写法都可以。 把包赋值给abcP，如果省略掉abcP，默认就是返回的包名【一般来讲我们都是写成和文件夹同名】（提倡省略的写法，除非本页中包存在同名的情况下）
	abcP "beegoStudy/abc" // 推荐使用这种写法，不要使用上面的写法
	_ "beegoStudy/routers" // _ ：代表这个包只执行里面的常量、变量和init函数，在本页面中并不使用（舍弃掉）
	"github.com/astaxie/beego"
	"beegoStudy/abc/hehe" // 特别注意了，这里返回的包名并不是hehe，而是hehehe。
	// 特别注意了：如果一个包被引入了多次，那么它就会初始化多少次（初始化包括：执行里面的常量、变量和init函数）
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"net/http"
	"html/template"
)

const MM string = "main的常量";

func main() {
    // beego.LoadAppConfig("ini", "conf/app2.conf") // 加载用户自定义配置文件，默认：conf/app.conf。调用多次，可加载多个配置文件，如果后面的文件和前面的 key 冲突，那么以最新加载的为最新值
	var str = "欧欧欧";
	abcP.Wl(str);
	// abc.wor(); // 这样调用是会出错的：cannot refer to unexported name abc.wor 小写字母开头的代表私有的函数，不能被外部调用（unexported）
	abcP.EchoMeile();
	abcP.EchoNiyehao();
	println(abcP.MM);
	println(MM);
	println(hehehe.MM);

	// beego.BConfig.WebConfig.Session.SessionOn = true // 使用session . 如果在配置文件中没有设置SessionOn的话。 如果都不设置的话是不能使用session的
	// 目前 session 模块支持的后端引擎包括 memory、cookie、file、mysql、redis、couchbase、memcache、postgres
	
	// 注册一个请求过滤函数，在路由之前
	// _ = beego.InsertFilter("/*", beego.BeforeRouter, FilterUserLogin); // 丢弃掉返回值，_= 也可以不写。 特别注意了，有返回值的函数不能写在函数外
	// beego.SetLogger("file", `{"filename":"/data/logs/go/beegoStudy.log"}`); // 设置日志文件路径，这样设置后，既会再控制台打印信息，又会把log写入文件
	beego.Debug("this is debug"); // 返回：2018/09/10 21:58:29.402 [D] [main.go:32] this is debug  （深蓝底色黑色字）
	beego.Alert("this is alert"); // 返回：2018/09/10 22:00:15.540 [A] [main.go:33] this is alert （浅蓝色）
    beego.Informational("this is informational"); // 返回：2018/09/10 22:10:16.161 [I] [main.go:34] this is informational （蓝色）
    beego.Emergency("this is emergency"); // 返回：2018/09/10 22:10:16.161 [M] [main.go:35] this is emergency （白色）
	beego.Critical("this is critical"); // 返回：2018/09/10 22:10:16.161 [C] [main.go:36] this is critical （紫色）
    beego.Error("this is error"); // 返回：2018/09/10 22:05:38.845 [E] [main.go:37] this is error （红色）
    beego.Warning("this is warning"); // 返回：2018/09/10 22:10:16.161 [W] [main.go:38] this is warning （黄色）
    beego.Notice("this is notice", " ou阳海雄", 35); // 返回：2018/09/10 22:06:22.349 [N] [main.go:39] this is notice （绿色）也可以打印多个变量
	// 上面是打印的日志，默认打印到控制台（console）
    // 如果只想输出到文件，就需要调用删除操作：beego.BeeLogger.DelLogger("console")
	// beego.SetLogFuncCall(true) 是否输出调用的文件名和文件行号；默认false:关闭，true：开启

    logs.SetLogger(logs.AdapterFile, `{"filename":"/data/logs/go/beegoStudy_logs.log", "level":7, "maxlines":0, "maxsize":2097152, "daily":true, "maxdays":20}`);
	// filename : 保存的文件路径
	// level：日志保存的级别，默认Trace级别
	// maxlines ： 每个文件保存的最大行数，默认值1000000
	// maxsize : 1<<21：2M 不能这样写，会报错的。每个文件保存的最大大小，默认1<<28:256MB
	// daily : 是否按照每天管理日记文件（logrotate），默认true
	// maxdays : 文件最多保存多少天，默认保存7天
	// rotate : 是否开启日记管理，默认true
    // perm : 日志文件权限
    logs.EnableFuncCallDepth(true); // logs日志默认输出调用的文件名和文件行号；默认true：开启，false：关闭
	logger();

    // beego.ErrorHandler("404", page_not_found); // 设置自定义404处理页面。注意后面的参数是函数名，并非字符串类型
	// beego.ErrorHandler("dbError", dbError);
	

	beego.Run()
}

// 下面演示添加过滤器
var FilterUserLogin = func (ctx *context.Context) { // 所有没有登录的请求都跳转到login页
	_, uid := ctx.Input.Session("uid").(int);
	if ((!uid) && ctx.Request.RequestURI != "/login") {
        ctx.Redirect(302, "/login");
	}
}

func logger() { // github.com/astaxie/beego/logs 的日记和 beego.XXX[Debug]的日记如出一辙，只是调用形式的不一样而已
    // an  official log.Logger
    logger := logs.GetLogger();
    logger.Println("this is a message of http"); // 2018/09/11 03:20:42.606 [main.go:71] this is a message of http  白色的字，相当于Emergency
	// an official log.Logger with prefix ORM
	logs.GetLogger("ORM").Println("this is a message of orm");  // 这个是带前缀的打印
    // 2018/09/11 03:26:24.703 [main.go:74] [ORM] this is a message of orm
    // 下面的打印都是默认自动换行的，且下面打印的行数是其调用函数所在的行数，这里为55行
    logs.Debug("my book is bought in the year of ", 2016); // 打印多个变量
    logs.Info("this %s cat is %v years old", "yellow", 3); // 替换多个变量
	logs.Warn("json is a type of kv like", map[string]int{"key":2016, "age":35}); // map类型也是可以直接打印的
	// 上面一句返回： 2018/09/11 03:08:54.910 [W] [main.go:55] json is a type of kv like map[key:2016 age:35]
    logs.Error(1024, "is a very", "good game"); // 打印多个变量
    logs.Critical("oh, crash"); // 严重的，（紫色）
}

func page_not_found(rw http.ResponseWriter, r *http.Request) { // 定义404错误处理页面
	// beego.Error(beego.BConfig.WebConfig.ViewsPath); // 返回：views
	t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html"); // 相对路径：views/404.html
	data := make(map[string]interface{});
	data["content"] = "page not found"
	t.Execute(rw, data);
}

func dbError(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("dberror.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/dberror.html");
	data := make(map[string]interface{});
	data["content"] = "database is now down";
	// beego.Info(data["content"]);
	t.Execute(rw, data);
}
