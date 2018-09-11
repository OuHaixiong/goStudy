package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"github.com/astaxie/beego/validation"
)
// 测试表单数据验证
type ValidateController struct {
    beego.Controller
}

type User struct {
	Name string
	Age int
}

func (c *ValidateController) Index() {
	u := User{"欧阳海雄", 40}
	// validationStruct := validation.Validation{} // 这种声明结构体的写法后面的{}不能少
	var validationStruct validation.Validation; // 声明一个结构体变量
	validationStruct.Required(u.Name, "name_required"); // name_required Can not be empty 
	validationStruct.MaxSize(u.Name, 2, "name_max"); // name_max Maximum size is 2
	validationStruct.Range(u.Age, 0, 18, "age") // 最后一个参数是指出错后对应的key值
    if validationStruct.HasErrors() { // 验证不通过后会继续下一个验证，如果单个字段有多个验证规则，全部都会进行验证
		for _, err := range validationStruct.Errors { // 如果有错误，打印错误信息
			beego.Error(err.Key, err.Message); // [E] [validate.go:27] age Range is 0 to 18
		}
	}

	// or use like this
    if v := validationStruct.Max(u.Age, 39, "age"); !v.Ok { // Max验证最大值，这里包括39
        beego.Error(v.Error.Key, v.Error.Message); // [E] [validate.go:33] age Maximum is 39 
	}
	// 定制错误信息
	minAge := 41 // Min ： 验证最小值（包括最小值）
	validationStruct.Min(u.Age, minAge, "age").Message("少儿不宜！"); // [E] [validate.go:41] age 少儿不宜！
	validationStruct.Min(u.Age, minAge, "age").Message("%d不禁", minAge); // 变量进行替换（ 错误信息格式化）
	if validationStruct.HasErrors() {
		for _, err := range validationStruct.Errors { 
			beego.Error(err.Key, err.Message);
		}
	}

	log.Println("验证完毕！！！"); // 不会有文件名和行号
	// c.Data["json"] = map[string]interface{}{"name":"欧海雄"};


	type my struct {
		Name string  // 结构体的属性如果大写的话，可以直接当json的数据输出
		age int      // 但是如果是小写，就不能输出了
	};
    myStruct := my{"欧阳海雄", 35};
    // c.Data["json"] = u; // 结构体是可以的，&u也是可以的
	c.Data["json"] = &myStruct; // 结构体声明在外面和在里面都是一样的
	beego.Informational(myStruct.age); // 2018/09/11 06:22:28.347 [I] [validate.go:58] 35 
    c.ServeJSON();
}

func (c *ValidateController) Get() {

}


