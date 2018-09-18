package controllers
// 测试错误处理，主要是自定义错误处理页面
import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found && this is error controller by Error404 func";
	c.TplName = "404.html";
}

func (c *ErrorController) Error501() {
	c.Data["content"] = "501   501  server error";
	// c.TplName = "501.tpl";
	c.TplName = "404.html"
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down && ErrorController->ErrorDb()";
	c.TplName = "dberror.html";
}
