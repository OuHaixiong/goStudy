// 测试绑定数据（binding数据）
// gin内置了几种数据的绑定例如JSON, XML等. 简单来说, 即根据Body数据类型, 将数据赋值到指定的结构体变量中. (类似于序列化和反序列化) 
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int `form:"age" json:"age"`
}

type Login struct {
	User     string `form:"user" json:"user" binding:"required"` // 注意:后面的form:user表示在form中这个字段是user,不是User, 同样json:user也是
	Password string `form:"password" json:"password" binding:"required"` // 注意:binding:"required"要求这个字段在client端发送的时候必须存在,否则报错!即参数必须存在
}



func main() {
	router := gin.Default()
	
	router.POST("/login", func (c *gin.Context) {
		var user User
		var err error
		contentType := c.Request.Header.Get("Content-Type")
        fmt.Println("Content-Type::::", contentType)
		switch contentType {
		    case "application/json" : err = c.BindJSON(&user)
		    case "application/x-www-form-urlencoded" : err = c.Bind(&user, binding.Form) // 默认使用form解析
		}
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"username" : user.Username,
			"password" : user.Password,
			"age" : user.Age,
		})
	})

	// 下面测试bind JSON数据
	router.POST("/bindJSON", funcBindJSON) // 测试请求详见goBase/http.go
	
	// 下面测试bind Form数据
    router.POST("/bindForm", funcBindForm)

	// 下面演示结构体转json字符串
    router.GET("/changeJson", funcChangeJson)

	// 下面演示返回xml
    router.GET("/someXML", funcSomeXML)
	
    router.Run(":8080");
}

// bind JSON 数据
func funcBindJSON(c *gin.Context) {
	var loginStruct Login
	if err := c.BindJSON(&loginStruct); err == nil { // .BindJSON本质是将request中的body中的数据按照json格式解析到loginStruct中
        if loginStruct.User == "ouhaixiong" && loginStruct.Password == "123456" {
			c.JSON(http.StatusOK, gin.H{"JSON==status":"you are logged in"})
			// 200 {"JSON==status":"you are logged in"}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"JSON==status":"unauthorized"})
			// 401 {"JSON==status":"unauthorized"}
		}
	} else {  //如果json格式错误的话，会执行到这里
		println("Error:::", err.Error()) // Error::: invalid character '}' looking for beginning of object key string
		c.JSON(404, gin.H{"JSON==>status":"binding JSON error!"});
		// 这里输出的http code 为 400 ， 而非上面的404，因为： [WARNING] Headers were already written. Wanted to override status code 400 with 404
        // {"JSON==\u003estatus":"binding JSON error!"}
	}
}

// bind form 数据
func funcBindForm(c *gin.Context) {
	var loginStruct Login
	// if err := c.BindWith(&loginStruct, binding.Form); err == nil { // .BindWith(&XXX, binding.Form)本质是坚request中的body数据解析到form中
    // 上面的方法：BindWith 是废弃的，提倡下面的写法：
	if err := c.ShouldBindWith(&loginStruct, binding.Form); err == nil { 
		// 也可以用 MustBindWith ，这个如果出错了，会直接返回 400的状态码。 而 ShouldBindWith 是可以自定义返回的状态码
		// c.Bind(&loginStruct) 默认使用form格式进行解析
		if loginStruct.User == "ouhaixiong" && loginStruct.Password == "123456" {
            c.JSON(http.StatusOK, gin.H{"Form==status":"you are logged in"})
		} else {
            c.JSON(http.StatusUnauthorized, gin.H{"Form==status":"unauthorized"})
		}
	} else { // 如果参数格式错误或者必传的参数没有传过来
		println("Error:::", err.Error())  // Error::: Key: 'Login.User' Error:Field validation for 'User' failed on the 'required' tag
		// Error::: Key: 'Login.Password' Error:Field validation for 'Password' failed on the 'required' tag
		c.JSON(500, gin.H{"Form==>status":"binding Form error!"})
		// 400 {"Form==\u003estatus":"binding Form error!"}
	}
}

func funcChangeJson(c *gin.Context) {
    var msg struct {
		Name    string `json:"user"` // 注意：在tag中指示的json显示为转为json字符串后对应的key（名称name）
		Message string `json:"mEss"`
		Number  int
	} 
	msg.Name = "OUouou"
	msg.Message = "海雄，你好"
	msg.Number = 123
	c.JSON(http.StatusOK, msg) // 框架自动进行了转换
	// 200 {"user":"OUouou","mEss":"海雄，你好","Number":123}
}

func funcSomeXML(c *gin.Context) { // 测试发送XML数据
	c.XML(http.StatusOK, gin.H{"username":"ouyanghaihai", "message":"HI hi Hi", "status":http.StatusOK});
	// <map><username>ouyanghaihai</username><message>HI hi Hi</message><status>200</status></map>
}

   
