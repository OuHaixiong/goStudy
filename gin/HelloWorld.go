package main

import (
   // "gopkg.in/gin-gonic/gin.v1" // 这个是老版本的写法，新版本中不能这样写了
    "github.com/gin-gonic/gin" // 引入框架包
    "net/http" // 引入http包
    "fmt"
)

func main() {
    port := ":8000";
    router := gin.Default();
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "<h3>Hello World</h3>"); // 输出字符串<h3>Hello World</h3>（并非html）
    });
    fmt.Println("服务已启动，监听端口为：", port);
    router.Run(port) // 监听的端口号。 到这来程序就不会再执行下去了，因为他一直处在监听端口的状态
}
