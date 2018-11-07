package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"time"
	"log"
	"io"
	"os"
)

func main() {
	router := gin.Default() // 默认是带有Logger和Recovery中间件的
	// router := gin.New() // 不带中间件的路由
	// router.Use(gin.Logger()) // 可以使用这种方式来指明中间件

	// 定义模板文件路径
    router.LoadHTMLGlob("templates/*") // 加载所有的模板文件（视图文件）（一定要这样写）

	/*router.GET("/user/:id", func(c *gin.Context) { // 路由的定义；后面的:id匹配任意字符
		id := c.Param("id"); // http://172.17.10.253:8000/user/ 这样是会报错的404
		c.String(http.StatusOK, "id=>%s", id);// 访问：http://172.17.10.253:8000/user/45
	});*/

    router.GET("/user/:name/*action", func(c *gin.Context) { // 这个路径和上面的冲突，路由不支持正则匹配
		name := c.Param("name"); // Context.Param("key") 获取url中匹配的参数
		action := c.Param("action") // 也可以这样写：c.Params.ByHName("XXX")
		message := name + " is " + action // http://172.17.10.253:8000/user/45/dsfdas/%E4%BD%A0%E5%A5%BD
		c.String(http.StatusOK, message) // 45 is /dsfdas/你好
		// : 和 * 的区别是，“:”只包含两个/内的字符，而“*”包括了前面的/和后面的所有字符
	})

    router.GET("/welcome", func(c *gin.Context) { // query 能获取 ? 后面的参数
		firstName := c.DefaultQuery("first_name", "Guest"); // 可以给默认值;特别需要注意的是：当first_name为空字串的时候，并不会使用默认的Guest值，空值也是值，DefaultQuery只作用于key不存在的时候，提供默认值。
		lastName := c.Query("last_name"); // 如果不传就是空字符串
		c.String(http.StatusOK, "Hello %s %s", firstName, lastName); // 经过urlencode编码
		// http://172.17.10.253:8000/welcome?last_name=%E6%B5%B7&first_name=ou  输出： Hello ou 海
		// 请求的url中最后是否带/都是一样的，如/welcome和/welcome/是一样的
	});

	// http的报文体(body)传输数据常见的格式就有四种。例如application/json，application/x-www-form-urlencoded, application/xml和multipart/form-data。
	// 后面一个主要用于图片上传。json格式的很好理解，urlencode其实也不难，无非就是把query string的内容，放到了body体里，同样也需要urlencode。
	// 默认情况下，c.PostFrom解析的是multipart/form-data和x-www-form-urlencoded的参数。
    router.POST("/form_post", func(c *gin.Context) { // 这个地址并不匹配/form_post/，这样post请求的话直接为空，并不报404，而如果是/form_post/d就会报404。而在get中后面的/是否有都是一样的
		message := c.PostForm("message");
		nick := c.DefaultPostForm("nick", "anonymous") // anonymous:无名的，假名的
		// 同样的道理如果nick传空字符串时，nick为空字符串，而并非是anonymous
		fmt.Println(nick);
		c.JSON(http.StatusOK, gin.H{ // c.String返回字符类型，c.JSON返回json数据
			"status" : gin.H{ // gin.H封装了生成json的方式，是一个强大的工具
				"status_code" : http.StatusOK,
				"status" : "ok",
			},
			"message" : message,
			"nick":nick,
		});
	});
	// curl -X POST http://127.0.0.1:8000/form_post -H "Content-Type:multipart/form-data" -d "message=hello&nick=欧阳海雄" 
	// 上面的返回：{"message":"","nick":"anonymous","status":{"status":"ok","status_code":200}}   因为multipart/form-data要传参数需要用: -F "xx=xxx"
	// curl -X POST http://127.0.0.1:8000/form_post -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nick=欧阳海雄"
	// 上面的返回：{"message":"hello","nick":"欧阳海雄","status":{"status":"ok","status_code":200}}
	// curl -X POST http://127.0.0.1:8000/form_post -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nick=欧阳海雄" | python -m json.tool
	// 返回的结果用python的工具进行格式化了，返回：
/* % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
	                              Dload  Upload   Total   Spent    Left  Speed
      100   115  100    84  100    31   4390   1620 --:--:-- --:--:-- --:--:--  6000
{
    "message": "hello",
    "nick": "\u6b27\u9633\u6d77\u96c4",
    "status": {
        "status": "ok",
        "status_code": 200
    }
} */


    loginEndpoint := func (c *gin.Context) {
        c.JSON(200, gin.H{
			"status" : 1,
			"message" : "It is OK!", // 注意这里的逗号（,）不能少，因为go语言默认是分号。不能写成 message:"xxx"
		})
        // c.Redirect(http.StatusMovedPermanently, "https://xxxx.xxx") // 重定向
	}
	submitEndpoint := func (c *gin.Context) {
        username := c.PostForm("username") // 如果不传，默认就是空字符串
        password := c.DefaultPostForm("password", "123456")
        c.JSON(http.StatusOK, gin.H{
			"status" : "-1",
			"message" : "This is post login.",
			"userInfo" : gin.H{
				"username" : username,
				"password" : password,
			}, // 注意这里的逗号（,）也不能少
		})
	}
	// 下面进行路由分组
	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.POST("/login", submitEndpoint) // 函数需要先声明，不让这里提示未找到该变量
		// curl -X POST http://172.17.10.253:8000/v1/login -H "Content-Type:application/x-www-form-urlencoded" -d "password=******&username=OuHaixiong"
		// 返回：{"message":"This is post login.","status":"-1","userInfo":{"password":"******","username":"OuHaixiong"}}
	}

    router.GET("/async", func (c *gin.Context) { // gin可以借助协程来实现异步任务，但是这时候得手动copy上下文，并且只能是可读取的
		cCopy := c.Copy() // 拷贝上下文
		go func () { // 声明一个匿名函数，并执行。 关键字go表示并发执行
			time.Sleep(5 * time.Second)
			log.Println("Done! in path " + cCopy.Request.URL.Path) // 这里会等待五秒后执行，打印：Done! in path /async
		}()
		// 这里的请求会直接返回，返回的为空
	})

	// 下面对gin进行单个文件上传测试 multipart/form-data
    router.POST("/upload", func(c *gin.Context) { // gin文件上传也很方便，和原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中了。
		name := c.PostForm("name") // 获取post上来的参数
		fmt.Println(name) // hai海
		file, header, err := c.Request.FormFile("upload") // 解析客户端文件name属性。如果不传文件，则会抛错，因此需要处理这个错误。一种方式是直接返回。
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request , error:", err.Error())
			return
		}
		filename := header.Filename // 获取上传过来的文件名
		fmt.Println(file, err, filename) // {0xc4201b2d50} <nil> productj.png
		
		out, err := os.Create(filename) // os.Create(filePath string) 通过文件路径，创建一个文件流操作
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close() // 延迟执行，即在function进行了return后会执行这个
		_, err = io.Copy(out, file) // io.Copy(target []byte, source []byte) 拷贝文件流数据
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusCreated, "upload successful") // http.StatusCreated = 201
		// curl -X POST http://172.17.10.253:8000/upload -H "Content-Type:multipart/form-data" -F "upload=@/data/www/productj.png" -F "name=hai海"
	})

	// 多文件上传
    router.POST("/multipart/upload", func(c *gin.Context) {
		err := c.Request.ParseMultipartForm(200000)
        if err != nil {
			log.Fatal(err)
		}
		formData := c.Request.MultipartForm // 得到文件句柄
		files := formData.File["upload"]
        for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				log.Fatal(err)
			}
			out, err := os.Create(files[i].Filename)
			defer out.Close()
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(out, file)
			if err != nil {
				log.Fatal(err)
			}
			c.String(http.StatusCreated, "Upload Successful!!!") // 如果有两个文件，打印：Upload Successful!!!Upload Successful!!!
		}
        // curl -X POST http://172.17.10.253:8000/multipart/upload -H "Content-Type:multipart/form-data" -F "upload=@/data/www/productj.png" -F "upload=@/data/www/125766383.jpg"
	})

	// 上传文件页面渲染
	router.GET("/upload", func(c *gin.Context) {
        c.HTML(http.StatusOK, "upload.html", gin.H{ // 渲染视图文件
			"singleUri" : "/upload", // 传递参数到视图页面
			"multipartUri" : "/multipart/upload",
		})
	})


	router.Run(":8000"); // 默认8080
}