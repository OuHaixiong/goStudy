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
	"time"
	// "io/ioutil"
)

func TestMiddleWare() gin.HandlerFunc { // 定义全局中间件
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("qq", "client_request") // 设置传递的参数。在路由中可以使用c.Get("qq")获得
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

func RunTime() gin.HandlerFunc { // 定义一个监控运行时间的中间件
    return func (c *gin.Context) {
		t := time.Now() // 获取当前时间，返回如：2018-12-13 01:25:45.933496744 -0500 EST m=+3.825548952
		fmt.Println("RunTime middleware :: now time :", t);
		c.Set("example", 123) // 设置example变量到Context的Key中,通过Get等函数可以取得
		c.Next()
		latency := time.Since(t) // 获取两个时间的差，返回微秒如：690.646µs
		log.Print("RunTime middleware :: ", latency) // 这个打印也会换行
		status := c.Writer.Status() // 这个c.Write是ResponseWriter,我们可以获得状态等信息
		log.Println("RunTime middleware :: ", status)
	}
}

func main() {
	router := gin.Default()

	router.Use(CORSMiddleware()) // 注册 CORS 跨域请求 中间件

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

    router.Use(RunTime()) // 注册多个全局中间件，执行顺序也是按注册顺序来的
	v2 := router.Group("/v2")
	// router.Use(RunTime()) // 写在下面是不行的，一定是需要在路由设置之前设置中间件
    v2.GET("/auth/signin", func(c *gin.Context) { // 这里只应用了全局中间件
        cookie := &http.Cookie{ // 申请cookie对象
			Name:     "session_id", // cookie 名
			Value:    "ouyanghaixiong", // cookie 值
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
		example := c.MustGet("example").(int) // 注意后面的.(int) 这里的类型要和前面中间件上c.Set的类型一致
		log.Print("++++ v2 -> home ++++")
		log.Println(example)
		c.JSON(http.StatusOK, gin.H{"controller":"/v2/home"})
	})

    v2.POST("/post-info", func (c *gin.Context) {
		fmt.Println("++++ v2 -> post-info ++++ start::::")
		fmt.Println("request method is:", c.Request.Method) // request method is: POST
		fmt.Println("request host is:", c.Request.Host) // request host is: 172.17.10.253:8080  包括了主机头和端口号
		fmt.Println("request url is:", c.Request.URL) // request url is: /v2/post-info?abc=abc&name=ouhaixiong&age=30 端口号后面的所有
		fmt.Println("request ContentLength is:", c.Request.ContentLength) // request ContentLength is: 239
		contentLength := c.Request.Header.Get("Content-Length")
		fmt.Printf("request Content-Length by header is %s \n", contentLength) // request Content-Length by header is 239
		c.Request.ParseForm() // 这句必不可少
		requestBodyData := c.Request.PostForm // application/x-www-form-urlencoded  POST请求获取form数据
		fmt.Printf("request body data is : %s \n", requestBodyData) // 如果是 application/x-www-form-urlencoded 请求的话，body内容如：request body data is : map[y:[yanayan] m:[铭铭]] 
		// requestBody, _ := ioutil.ReadAll(c.Request.Body) // multipart/form-data  POST请求获取body数据
		// requestBodyData := string(requestBody)
		// fmt.Printf("request body data is : %s \n", requestBodyData) // 如果是 multipart/form-data 请求的话，body内容如下：
		/*------------------------------c32bbf69d467
		Content-Disposition: form-data; name="y"
		
		yanayan
		------------------------------c32bbf69d467
		Content-Disposition: form-data; name="m"
		
		铭铭
		------------------------------c32bbf69d467--
		 */
		requestQueryAll := c.Request.URL.Query() // 获取所有url请求参数；返回map结构
		fmt.Printf("request url data is : %s \n", requestQueryAll) // request url data is : map[age:[30] abc:[abc] name:[ouhaixiong]] 
		fmt.Println("++++ v2 -> post-info ++++ end::::")
	})
	// curl -X POST --cookie "session_id=ouyanghaixiong" "http://172.17.10.253:8080/v2/post-info?abc=abc&name=ouhaixiong&age=30" -H "Content-Type:multipart/form-data" -F "y=yanayan" -F "m=铭铭"

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

// CORS middleware （CORS 跨域请求中间件）
func CORSMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
		log.Print("CORS middleware")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Yf-Country")
		if c.Request.Method == "OPTIONS" {
			// c.Abort(200) // 不能这样写，报错：too many arguments in call to c.Abort
			//c.Abort() // Abort 固定返回404
            c.String(http.StatusOK, "OK")
			log.Print("CORS middleware :: 200")
			return // 就算这里全局返回了，后面的所有全局中间件还是会执行一遍，只是路由就不会执行
		}
		c.Next()
	}
}
