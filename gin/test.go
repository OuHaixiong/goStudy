package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main() {
	router := gin.Default()
	/*router.GET("/user/:id", func(c *gin.Context) { // 路由的定义
		id := c.Param("id"); // http://172.17.10.253:8000/user/ 这样是会报错的404
		c.String(http.StatusOK, "id=>%s", id);// 访问：http://172.17.10.253:8000/user/45
	});*/

    router.GET("/user/:name/*action", func(c *gin.Context) { // 这个路径和上面的冲突，路由不支持正则匹配
		name := c.Param("name");
		action := c.Param("action")
		message := name + " is " + action // http://172.17.10.253:8000/user/45/dsfdas/%E4%BD%A0%E5%A5%BD
		c.String(http.StatusOK, message) // 45 is /dsfdas/你好
	})

    router.GET("/welcome", func(c *gin.Context) { // query 能获取 ? 后面的参数
		firstName := c.DefaultQuery("first_name", "Guest"); // 可以给默认值;特别需要注意的是：当firstname为空字串的时候，并不会使用默认的Guest值，空值也是值，DefaultQuery只作用于key不存在的时候，提供默认值。
		lastName := c.Query("last_name"); // 如果不传就是空字符串
		c.String(http.StatusOK, "Hello %s %s", firstName, lastName); // 经过urlencode编码
		// http://172.17.10.253:8000/welcome?last_name=%E6%B5%B7&first_name=ou  输出： Hello ou 海
	});

	// http的报文体(body)传输数据常见的格式就有四种。例如application/json，application/x-www-form-urlencoded, application/xml和multipart/form-data。
	// 后面一个主要用于图片上传。json格式的很好理解，urlencode其实也不难，无非就是把query string的内容，放到了body体里，同样也需要urlencode。
	// 默认情况下，c.PostFROM解析的是x-www-form-urlencoded或from-data的参数。
    router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message");
		nick := c.DefaultPostForm("nick", "anonymous") // anonymous:无名的，假名的
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
	// 上面的返回：{"message":"","nick":"anonymous","status":{"status":"ok","status_code":200}}
	// curl -X POST http://127.0.0.1:8000/form_post -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nick=欧.海雄"
	// 上面的返回：{"message":"hello","nick":"欧阳海雄","status":{"status":"ok","status_code":200}}
	// curl -X POST http://127.0.0.1:8000/form_post -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nick=欧阳海雄" | python -m json.tool
	// 格式化了，返回：
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
	router.Run(":8000"); 
}