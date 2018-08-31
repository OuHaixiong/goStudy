package routers

import (
	"beegoStudy/controllers"
	"github.com/astaxie/beego"
	"beegoStudy/controllers/adminUser"
)

func init() {
	beego.Router("/", &controllers.MainController{}) // 路由注册（映射 URL 到 controller） 。 
	// &controllers.MainController{} : 声明一个controllers包下面的MainController结构体变量并初始化成员属性（这里属性为空），最后按引用传递

	beego.Router("/admin-user", &controllers_adminUser.UserController{}, "get:List") // 第三个参数为注册自定义函数来处理请求，形式：method:function,可以为： *：func
    beego.Router("/admin-user/user/test", &controllers_adminUser.UserController{}, "get:Test")
}
