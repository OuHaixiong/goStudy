package main
 
import (
    "fmt"
    "net/http"
    "reflect"
    "strings"
    "net/url"
    "html"
)
 
func hello(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello"))
    uri := req.URL.String() // /hello?dd=dd&abc=123 返回端口号往后的所有url，不包括锚
    fmt.Println(uri)
    // curl -X GET "http://127.0.0.1:8804/hello?dd=dd&abc=123#mao"
    // url编码
    str := "中 文-_.编abc码"
    str = "abc=" + str // 连接两个字符串，如果是大量大文本拼接，用 bytes.Buffer
    str = url.QueryEscape(str) // 对url进行encode (和php的urlencode是一样的)
    fmt.Printf("url.QueryEscape: %s \n", str)  // url.QueryEscape: abc%3D%E4%B8%AD+%E6%96%87-_.%E7%BC%96abc%E7%A0%81 
    s, _ := url.QueryUnescape(str) // 对url进行decode （和php的urldecode是一样的）
    fmt.Printf("url.QueryUnescape: %s \n", s) // url.QueryUnescape: abc=中 文-_.编abc码

    hstr := "< >'\""
    hstr = html.EscapeString(hstr) // 对特殊字符进行html实体字符转换（字符转码和php的htmlspecialchars）。 注意空格并不会进行实体转义
    fmt.Printf("html.EscapeString: %s \n", hstr) // html.EscapeString: &lt; &gt;&#39;&#34;
    hstr = html.UnescapeString(hstr) // 将html实体字符进行解析（和php的htmlspecialchars_decode一样）
    fmt.Printf("html.UnescapeString: %s \n", hstr) // html.UnescapeString: < >'"
}
 
type Handlers struct {
}
 
func (h *Handlers) ResAction(w http.ResponseWriter, req *http.Request) {
    fmt.Println("res")
    w.Write([]byte("res"))
}
 
func say(w http.ResponseWriter, req *http.Request) {
    pathInfo := strings.Trim(req.URL.Path, "/")
    fmt.Println("pathInfo:", pathInfo)
 
    parts := strings.Split(pathInfo, "/")
    fmt.Println("parts:", parts)
 
    var action = "ResAction"
    fmt.Println(strings.Join(parts, "|")) // 通过一个字符串，连接切片中的值；第一个参数一定是 []string
    if len(parts) > 1 {
        fmt.Println("22222222")
        action = strings.Title(parts[1]) + "Action"
    }
    fmt.Println("action:", action)
    handle := &Handlers{}
    controller := reflect.ValueOf(handle)
    method := controller.MethodByName(action)
    r := reflect.ValueOf(req)
    wr := reflect.ValueOf(w)
    method.Call([]reflect.Value{wr, r})
}
 
func main() { // 最简单的http服务；访问：ip:8804/handle/res
    http.HandleFunc("/hello", hello)
    http.Handle("/handle/", http.HandlerFunc(say))
    var port = ":8804"
    fmt.Println("listen is ", port)
    http.ListenAndServe(port, nil)
    //select {} //阻塞进程
}
