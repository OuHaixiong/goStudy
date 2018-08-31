package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller // 内嵌了 beego.Controller，这就是 Go 的嵌入方式，也就是 MainController 自动拥有了所有 beego.Controller 的方法。(类似继承)
	// beego.Controller 拥有很多方法，其中包括 Init、Prepare、Post、Get、Delete、Head 等方法
}

func (c *MainController) Get() { // 重写Get方法
	c.Data["Website"] = "beego.me" //赋值到模板文件中，模板文件可直接通过：{{.Website}} 获取对应的值
	c.Data["Email"] = "258333309@qq.com"
	c.Data["name"] = "欧阳海雄";
	c.TplName = "index.tpl" // 设置模板文件的路径，这里的路径是相对views的。 默认支持tpl 和 html 的后缀名
	// 默认是会调用模板文件进行渲染，也可以不使用模板文件，直接输出字符串，如：
	// c.Ctx.WriteString("hello"); // controller.Ctx.WriteString() 输出字符串
}
