package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"github.com/astaxie/beego/validation"
	"strings"
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
		for _, err := range validationStruct.Errors { // 这里会把之前的和现在的全部都打印（相当于调用方法就进行了验证）
			beego.Error(err.Key, err.Message); // [E] [validate.go:41] age 少儿不宜！
		}
	}

	log.Println("验证完毕！！！"); // 不会有文件名和行号
	// c.Data["json"] = map[string]interface{}{"name":"欧海雄"}; // map、结构体均可直接当json输出
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
	validationStruct := validation.Validation{};
	m := member{
		Name   : "Bear-Ou", // Name.Match Must match ^Bear.* 
		Age    : 18, // Age.Range Range is 1 to 140 
		Email  : "Bear@maimengmei.com", // 这样写的话，后面的“,”不能省略，因为go一行结束后如果不写，默认为“;”；所以会报错
		Mobile : "+8615019261350", // Mobile.Mobile Must be valid mobile number  (+8615019261350是可以的)
		Ip     : "12.86.95.30", // Ip.IP Must be a valid ip address 
	}
	boolean, err := validationStruct.Valid(&m);
	if err != nil { // handle error
        beego.Error("验证结构体数据出错了：", err);
	}
	if (!boolean) { // validation does not pass
        for _, err := range validationStruct.Errors { // 如果验证不通过，打印所有出错信息
			beego.Error(err.Key, err.Message);
		}
	}
    c.Data["json"] = m; // 赋值时，带“&”和不带均可
	// memberStruct := member{};
    // c.Data["json"] = &memberStruct; // 这样打印出来的都是0和""
    c.ServeJSON(); // 输出json格式的数据
}

// 验证函数写在“valid” tag的标签里。
// 各个函数之间用分号“;”分隔，分号后面可以有空格
// 参数用括号“()”括起来，多个参数之间用逗号“,”分开，逗号后面可以有空格。
// 正则函数（Match）的匹配模式用两斜杆“//”括起来
// 各个函数的结果的key值为字段名。
type member struct {
	Id     int
	Name   string `valid:"Required; Match(/^Bear.*/)"` // Name不能为空并且必须以Bear开头
	Age    int    `valid:"Range(1, 140)"` // Age必须是1~140，超出此范围即为不合法
	Email  string `valid:"Email;MaxSize(100)"` // Email必须符合邮箱格式，并且最大长度不能大于100个字符
	Mobile string `valid:"Mobile"` // Mobile必须为正确的手机号
    Ip     string `valid:"IP"` // Ip必须为一个正确的IPv4地址
}

// 如果你的结构体（struct）实现了接口validation.ValidFormer
// 当Struct Tag中的测试都成功时，将会执行Valid(接口函数)进行自定义验证
func (m *member) Valid(v *validation.Validation) {
    if strings.Index(m.Name, "admin") != -1 {
        v.SetError("Name", "名称里不能包含“admin”"); //  通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
    }
}

/*
Struct Tag 可用的验证函数：

Required 不为空，即各个类型要求不为其零值
Min(min int) 最小值，有效类型：int，其他类型都将不能通过验证
Max(max int) 最大值，有效类型：int，其他类型都将不能通过验证
Range(min, max int) 数值的范围，有效类型：int，他类型都将不能通过验证
MinSize(min int) 最小长度，有效类型：string slice，其他类型都将不能通过验证
MaxSize(max int) 最大长度，有效类型：string slice，其他类型都将不能通过验证
Length(length int) 指定长度，有效类型：string slice，其他类型都将不能通过验证
Alpha alpha字符，有效类型：string，其他类型都将不能通过验证
Numeric 数字，有效类型：string，其他类型都将不能通过验证
AlphaNumeric alpha 字符或数字，有效类型：string，其他类型都将不能通过验证
Match(pattern string) 正则匹配，有效类型：string，其他类型都将被转成字符串再匹配(fmt.Sprintf(“%v”, obj).Match)
AlphaDash alpha 字符或数字或横杠 -_，有效类型：string，其他类型都将不能通过验证
Email 邮箱格式，有效类型：string，其他类型都将不能通过验证
IP IP 格式，目前只支持 IPv4 格式验证，有效类型：string，其他类型都将不能通过验证
Base64 base64 编码，有效类型：string，其他类型都将不能通过验证
Mobile 手机号，有效类型：string，其他类型都将不能通过验证
Tel 固定电话号，有效类型：string，其他类型都将不能通过验证
Phone 手机号或固定电话号，有效类型：string，其他类型都将不能通过验证
ZipCode 邮政编码，有效类型：string，其他类型都将不能通过验证
*/