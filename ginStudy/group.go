// 测试路由分组，测试中间件
package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"strings"
	"os"
	"path/filepath"
)

func TestMiddleWare() gin.HandlerFunc { // 定义全局中间件
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("qq", "client_request") // 设置传递的参数
		c.Next() // 中间件和路由的处理流程是，先运行中间件，直到遇到c.Next()后会运行路由的代码，完了后会接着运行c.Next()之后的代码
		fmt.Println("after middleware")
	}
}

func SingleMiddleWare() gin.HandlerFunc { // 定义单个中间件
	return func(c *gin.Context) {
		fmt.Println("单个中间件开始了")
		c.Next()
		fmt.Println("单个中间件结束了")
	}
}

func GroupMiddleWare() gin.HandlerFunc { // 定义群组中间件
	return func(c *gin.Context) {
		fmt.Println("group middler ware start++++++++")
		c.Next()
		fmt.Println("group middler ware end +++++++++")
	}
}

func AuthMiddleWare() gin.HandlerFunc { // 定义中间件，模拟验证过程
    return func(c *gin.Context) {
		fmt.Println("auth middle ware ::::::: start")
        if cookie, err := c.Request.Cookie("session_id"); err == nil {
			// c.Request.Cookie("XXX") 获取cookie
			value := cookie.Value
			fmt.Println(value)
			if value == "ouyanghaixiong" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{ // http.StatusUnauthorized = 401
			"error" : "Unauthorized",
		})
		c.Abort() // 终止请求。 终止请求后，路由中的内容不会再执行，但是之前中间件的内容依然还是会执行
		fmt.Println("auth middle ware ::::::: end")
		return
	}
}


func main() {
	router := gin.Default()

    router.GET("/single/test", SingleMiddleWare(), func(c *gin.Context) {
		fmt.Println("这里是单个控制器开始了：/single/test")
		c.JSON(http.StatusOK, gin.H{
			"middle_ware": "SingleMiddleWare",
		})
	})
	// curl -X GET http://172.17.10.253:8080/single/test

	router.Use(TestMiddleWare()) // 注册全局中间件
	// { 这里的括号可以不要。只要注册了全局中间件，后面定义的路由都会执行全局中间件
    // 需要注意，虽然名为全局中间件，只要注册中间件的过程之前设置的路由，将不会受注册的中间件所影响。只有注册了中间件以下代码的路由函数规则，才会被中间件装饰。
		router.GET("/bb", func(c *gin.Context) {
			fmt.Println("handle bb request-----")
			request := c.MustGet("qq").(string)
			req, _ := c.Get("qq") // Get和MustGet 这两个方法都是获取中间件中设置的参数值；区别在于如果没有注册就使用MustGet方法读取c的值将会抛错，可以使用Get方法取而代之
			// 即如果在中间件中没有qq这个变量，使用MustGet会直接报错，而Get不会
			c.JSON(http.StatusOK, gin.H{
				"request" : request, // request	"client_request"
				"req" : req,         // req	    "client_request"
			})
		})
	// }  // 使用花括号包含被装饰的路由函数只是一个代码规范，即使没有被包含在内的路由函数，只要使用router进行路由，都等于被装饰了。
	// curl -X GET http://172.17.10.253:8080/bb

	router.GET("/single/all", SingleMiddleWare(), func(c *gin.Context) {
		// 多个中间件执行的顺序是 全局=>单个, 比如这里返回：
        // before middleware
        // 单个中间件开始了
        // 这里是单个控制器开始了：/single/all
        // 单个中间件结束了
        // after middleware
		fmt.Println("这里是单个控制器开始了：/single/all")
		c.JSON(http.StatusOK, gin.H{
			"controller": "/single/all",
			"middleWare": "SingleMiddleWare",
		})
	})
	// curl -X GET http://172.17.10.253:8080/single/all

	hehe := func (c *gin.Context) {
		fmt.Println("------v1->hehe------")
        c.JSON(200, gin.H{
			"status" : 1,
			"message" : "v1 hehe",
		})
	}
	v1 := router.Group("/v1")  // 为了方便前缀相同的URL的管理
	// v1 := router.Group("/", GroupMiddleWare()) // 群组中间件，也可以下面的写法
	v1.Use(GroupMiddleWare()) // 为路由分组单独指定中间件
	// 特别注意了：已经设置过的全局中间件，对后面的所有路由均有作用，同样分组路由也不例外
	{ // 这个括号，可以写也可以不写
		v1.GET("/login", func(c *gin.Context) {
			fmt.Println("===v1->login===");
			c.String(http.StatusOK, "v1 login")
		})
		// curl -X GET http://127.0.0.1:8080/v1/login
		v1.GET("/hehe", hehe)
		// curl -X GET http://127.0.0.1:8080/v1/hehe
    }

	v2 := router.Group("/v2")
    v2.GET("/auth/signin", func(c *gin.Context) { // 这里只应用了全局中间件
        cookie := &http.Cookie{ // 申请cookie对象
			Name:     "session_id", // cookie 名
			Value:    "ouyanghaixiong8", // cookie 值
			Path:     "/", // 保存路径
			HttpOnly: true, // 是否只读
		}
		http.SetCookie(c.Writer, cookie) // 在浏览器中设置cookie
		c.String(http.StatusOK, "Login successful") // http.StatusOK = 200
	})
	v2.Use(GroupMiddleWare(), AuthMiddleWare()) // 全局中间件 => 多个中间件，中间件按书写顺序依次嵌套执行
	v2.GET("/login", func(c *gin.Context) {
		fmt.Println("++++ v2 -> login ++++")
		c.String(http.StatusOK, "v2 login")
	})
	v2.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"controller":"/v2/home"})
	})

	// 下面测试静态文件服务
	// fmt.Println(getCurrentDirectory())
    router.StaticFS("/show-dir", http.Dir("./showDir")) // 显示当前目录showDir下的所有文件和文件夹
	// router.Static("/files", "/bin") // 这句无法执行成功，报404
    router.StaticFile("/test-img", "./showDir/productj.png") // 显示指定文件（静态资源，不需要打包即可更新）


	router.Run() // 默认8080
	// 还可以这样写  http.ListenAndServe(":8080", router)
	// 也可以如下写法：
	// s := &http.Server{
    //     Addr:           ":8000",
    //     Handler:        router,
    //     ReadTimeout:    10 * time.Second,
    //     WriteTimeout:   10 * time.Second,
    //     MaxHeaderBytes: 1 << 20,
    // }
    // s.ListenAndServe()

}

// 这里返回：/tmp/go-build057865520/b001/exe
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) // 返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err.Error())
	}
	return strings.Replace(dir, "\\", "/", -1) // 将\替换成/
}
