package controllers_adminUser // 包名不能包含中划线（-）和斜线（/）

import (
	"github.com/astaxie/beego"
	// ""
)

type UserController struct {
	beego.Controller
}

func (c *UserController) List() {
	println("controllers/adminUser -> List action")
	c.TplName = "usercontroller/get.tpl" // 设置视图路径，这里直接是views下的路径，不用写views
}

func (c *UserController) Get() {
	println("controllers/adminUser -> List action")
	// 如果没有指定渲染视图，那么就会按照默认的视图进行渲染：controller name/method name.tpl 。 在这里默认视图为views/usercontroller/get.tpl
}

func (c *UserController) Test() {
	appName := beego.AppConfig.String("appname") // 获取配置项。如果获取一个并不存在的key时，或直接返回空字符串
	runMode := beego.AppConfig.String("runmode");
	mysqlPassword := beego.AppConfig.String("mysqlpass");

	c.Ctx.WriteString(appName); // 输出字符串，html是不支持的
	c.Ctx.WriteString(" " + mysqlPassword + " "); // 输出只能有一个参数
	c.Ctx.WriteString(runMode);
	beego.AppConfig.Set("appname", "无聊") // 通过程序代码修改配置项
	c.Ctx.WriteString(beego.AppConfig.String("appname"));
	c.Ctx.WriteString(beego.AppConfig.String("redis::mysqlpass")); // 获取段（section）中的配置
	// c.Ctx.WriteString(beego.GetConfig("test", "logpath")); // 这写法不对
	c.Ctx.WriteString(beego.AppConfig.String("mongodb::username"));

	beego.BConfig.ServerName = "beego" // 通过beego.BConfig.xxx也可以获取和设置默认的配置项
	c.Ctx.WriteString(beego.BConfig.ServerName);
}