package routers

import (
	"beegoStudy/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}) // 路由注册（映射 URL 到 controller） 。 
	// &controllers.MainController{} : 声明一个controllers包下面的MainController结构体变量并初始化成员属性（这里属性为空），最后按引用传递
}
