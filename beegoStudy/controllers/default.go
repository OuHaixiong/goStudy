package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "258333309@qq.com"
	c.Data["name"] = "欧阳海雄";
	c.TplName = "index.tpl"
}
