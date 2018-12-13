// 测试go的http请求
package main

import (
	"net/http"
	"io/ioutil"
	"strings"
	// "os"
	// "io"
	// "bytes"
	// "mime/multipart"
)

func main() {
	// response, _ := http.Get("http://172.17.10.253:8000/upload") // 调用最基本的GET,并获得返回值
	// body, err := ioutil.ReadAll(response.Body)
	// if err == nil {
	// 	println(string(body))
	// }


	// response, err = http.Post("http://172.17.10.253:8000/v1/login", "application/x-www-form-urlencoded", strings.NewReader("username=ouhaihai&password=****")) // 模拟POST请求
	// defer response.Body.Close()
	// body, err = ioutil.ReadAll(response.Body)
	// if err == nil {
	// 	println(string(body))
	// }


	// 下面演示文件上传(form post请求)
	/*buffer := new(bytes.Buffer) // 声明一个buffer（缓冲区）对象
	writer := multipart.NewWriter(buffer) // 通过buffer声明一个写缓冲区对象
	formFile, _ := writer.CreateFormFile("upload", "abc.jpeg") // 创建一个表单文件对象，设置名称（name）和文件名（filename）
	fileOpen, _ := os.Open("/data/www/125766383.jpg") // 通过os打开一个读取文件流对象
	defer fileOpen.Close()
	io.Copy(formFile, fileOpen) // 拷贝文件流
	writer.Close()
	response, err = http.Post("http://172.17.10.253:8000/upload", writer.FormDataContentType(), buffer) // writer.FormDataContentType() = multipart/form-data
	if err != nil {
		println(err)
	}
	readResponse(response)*/


    // 下面模拟json格式的post请求
	// response, err := http.Post("http://127.0.0.1:8080/bindJSON", "application/json", strings.NewReader("{\"user\":\"ouhaixiong\", \"password\":\"123456\"}"))
	// response, err := http.Post("http://127.0.0.1:8080/bindJSON", "application/json", strings.NewReader("{\"user\":\"ouhaixiong\", \"password\":123}"))
	// response, err := http.Post("http://127.0.0.1:8080/bindJSON", "application/json", strings.NewReader("{\"user\":\"ouhaixiong\", \"password\":\"12345d6\"}"))
	// if err != nil {
	// 	println("Error::::", err)
	// }
	// readResponse(response)

	// 下面模拟form表单post数据
	response, err := http.Post("http://127.0.0.1:8080/bindForm", "application/x-www-form-urlencoded", strings.NewReader("user=ouhaixiong&password=123456&ddd&ddt="))
    if (err != nil) {
		println("Error::::", err.Error()); // Error:::: Post http://127.0.0.1:8080/bindForm: dial tcp 127.0.0.1:8080: connect: connection refused
		return
	}
	readResponse(response)

	response, err = http.Get("http://127.0.0.1:8080/changeJson");
	readResponse(response);

	response, _ = http.Get("http://127.0.0.1:8080/someXML")
	readResponse(response);
	

	
}

func readResponse(response *http.Response) { // 用于读取response中的body
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		println(response.StatusCode)
        println(string(body))
	} else {
        println("Error:", err)
	}
}
