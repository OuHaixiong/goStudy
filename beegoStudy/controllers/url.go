package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type UrlController struct {
	beego.Controller
}

func (c *UrlController) Index() {
	c.Data["username"] = "欧阳海雄"
	// c.Data["PersonUrl"] = c.UrlFor("UrlController.Get", ":last", "hai", ":first", "ou"); // 这也找不到，操！版本升级了，方法名变了，变成了URLFor
	c.Data["PersonUrl"] = c.URLFor("UrlController.Get", ":last", "hai", ":first", "ou");
	ext := c.Ctx.Input.Param(":ext")  // ?:ext=888 问号后面的参数并不会直接传入进来
	fmt.Printf("%T \n", ext) // 返回string
	c.Data["Ext"] = ext
	params := c.Ctx.Input.Params()
	// c.Data["p1"] = c.Ctx.Input.Params["1"] // 这样的写法是错误的
	for k, v := range params { // 除了前缀两个 /:controller/:method 的匹配之外，剩下的 url beego 会帮你自动化解析为参数，保存在 this.Ctx.Input.Params 当中
		println(k, "=>", v) // /url/index/adsdasD/ou/hai?:ext=888
// :splat => adsdasD/ou/hai
// 0 => adsdasD
// 1 => ou
// 2 => hai
	}
	c.Data["Params"] = params
}

func (c *UrlController) Get() {
	c.Data["Username"] = "ouhaixiong";
	c.Ctx.Output.Body([]byte("ok"))
}

func (c *UrlController) List() {
    c.Ctx.Output.Body([]byte("I am list")) // []byte(string):字符串转二进制。 controller.Ctx.Output.Body():输出二进制数据，在浏览器中会打印：I am list
}

func (c *UrlController) Myext() {
	c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":ext")))
}

func (c *UrlController) GetUrl() {
    c.Ctx.Output.Body([]byte(c.URLFor(".Myext"))); // 打印：/url/myext
}


