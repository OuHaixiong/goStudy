package controllers

import (
	"github.com/astaxie/beego"

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
