package routers

import (
	"beegoStudy/controllers"
	"github.com/astaxie/beego"
	"beegoStudy/controllers/adminUser"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{}) // 路由注册（映射 URL 到 controller） 。 
	// &controllers.MainController{} : 声明一个controllers包下面的MainController结构体变量并初始化成员属性（这里属性为空），最后按引用传递

    // 下面的形式叫固定路由
	beego.Router("/admin-user", &controllers_adminUser.UserController{}, "get:List") // 第三个参数为注册自定义函数来处理请求，形式：httpmethod:functionname,可以为： *：func  *表示任意的 method 都执行该函数
	// 如果是同一个url地址，不同的请求方式执行不同的函数，可以用;分隔；如：beego.Router("/simple", &controllers.SimpleController{}, "get:GetFunc;post:PostFunc");   还可以："get,post:ApiFunc"
	beego.Router("/admin-user/user/test", &controllers_adminUser.UserController{}, "get:Test")
	
	// 最基本的路由
	beego.Get("/hello", func (ctx *context.Context) { // 如果是post请求：.Post(...)...
        ctx.Output.Body([]byte("hello world我爱你")) // 输出二进制流
	});

	beego.Any("/foo", func (ctx *context.Context) { // 注册一个可以响应任何 HTTP 的路由；可post、get、put请求等
		ctx.Output.Body([]byte("bar foo Bear-欧阳海雄"))
	})
	// 其他支持的基础函数有：beego.Put(router, beego.FilterFunc) 、Patch、Head、Options、Delete

	// 下面演示正则路由
	// /?:xx 表示匹配0次或一次
	// beego.Router("/test/api/?:id", &controllers.TestController{}, "get:Api"); // 大小写匹配。这个路由匹配：/test/api/80 或 /test/api 或 /test/api/?dd=xxx
	// beego.Router("/test/api/:id", &controllers.TestController{}, "get:Api"); // 仅匹配：/test/api/80，不匹配：/test/api 和 /test/api/ 。 /:xx表示精准匹配
	// beego.Router("/test/api/:id([0-9]+)", &controllers.TestController{}, "get:Api"); // 匹配：/test/api/007 ， 不匹配：/test/api/wuyu
	// beego.Router("/test/api/:id([\\w]+)", &controllers.TestController{}, "get:Api"); // 即匹配 /test/api/007 又匹配 /test/api/wuyu ，但不匹配： /test/api/无聊
    // beego.Router("/test/download/*.*", &controllers.TestController{}, "get:Download"); // 匹配：/test/download/file/api.xml?name=Bear ；此时，:path为file/api，:ext为xml
	beego.Router("/test/download/*", &controllers.TestController{}, "get:Download"); // 全匹配方式。 匹配：/test/download/file/api.json；此时，:splat为file/api.json
	// beego.Router("/test/api/:id:int", &controllers.TestController{}, "get:Api"); // int 类型设置方式，:int为匹配:id为int类型，框架帮你实现了正则 ([0-9]+)；和正则:id([0-9]+)是一样的
	// beego.Router("/test/api/:id:string", &controllers.TestController{}, "get:Api"); // :string 和 ([\\w]+) 是一样的
	beego.Router("test/api/bear_:id([\\d]+).html", &controllers.TestController{}, "get:Api"); // 带前缀的自定义正则也是可以的
	

	// 下面演示自动路由
	beego.AutoRouter(&controllers.TestController{}); // 这样如果访问：/test/index 就会调用TestController的Index方法。请求的url中的方法名都会转换为小写，并且所有的后缀都可以匹配到：.html、.xml、.json
	beego.Include(&controllers.TestController{}) // 使用注解的方式来注册路由。效果和自己通过 Router 函数注册是一样的
	// 注意注解的方法注册路由仅在dev模式下才能进行生成，生成的路由放在/routers/commentsRouter_controllers.go文件中

	// 下面的路由是演示URL构建用的
	beego.Router("/wu/liao", &controllers.UrlController{}, "*:List")
	beego.Router("/person/:last/:first", &controllers.UrlController{}) // 如果是get请求/person/ddd/aaa ， 将交给UrlController->Get()处理
	beego.AutoRouter(&controllers.UrlController{}); // get:/url/index => UrlController->index()

	beego.Router("/login", &controllers.TestController{}, "GET:Login") // 最后这个参数的method大小写均可。  登录页

    beego.AutoRouter(&controllers.ValidateController{});
}
