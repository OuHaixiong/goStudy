// 文件操作
package main

import (
	"fmt"
	// "crypto/tls"
	"io/ioutil"
	// "net/http"
	"log"
	"os"
	"io"
	"bufio"
	"path/filepath"
	"time"
)

func check(e error) {
    if (e != nil) {
		panic(e) // 如果失败返回如下信息：
		/*panic: open test.txt: permission denied

goroutine 1 [running]:
main.check(0x4d90c0, 0xc42008c1e0)
	/data/www/goStudy/file.go:70 +0x4a
main.main()
	/data/www/goStudy/file.go:64 +0x44d
exit status 2
*/
	}
}

func main() {
	/*
	link := "http://xkcd.com/55"
	tr := &http.Transport{
		TLSClientConfig : &tls.Config{InsecureSkipVerify:true}, // 特别注意了：最后这个“,”不能省略，不然包语法错误：syntax error: unexpected newline, expecting comma or }
	}
	client := &http.Client{Transport:tr} // 声明一个http请求客户端
	response, err := client.Get(link) // get请求访问一个地址
	if (err != nil) {
		log.Fatal(err)
	}
	defer response.Body.Close() // defer 延迟执行，在函数执行完后（return）再执行。defer执行顺序为先进后出
	// block forever at the next line
	content, err := ioutil.ReadAll(response.Body) // 读取二进制流文件。 func ReadAll(r io.Reader) ([]byte, error) 读取r中的所有数据，并返回读取的数据和遇到的错误
	if err != nil {
        fmt.Println("read body is error:", err)
	}
	fmt.Println(string(content)) // 二进制流转为字符串输出：<!DOCTYPE html> <html>...
	// string([]byte) 二进制流转字符串
	fmt.Println(content) // 输出二进制文件：[60 33 68 79 67 84 89 80 69 32...
	*/
	
	files, err := ioutil.ReadDir("/data/www/") // func ReadDir(dirname string) ([]os.FileInfo, error) 读取指定目录中的所有目录和文件（不包括子目录）.dirname后面有无/都一样
	// ReadDir 返回读取到的文件信息列表和遇到的错误，列表是经过排序的（按正序排列：数字0...1，A->Z,a->z）
	if (err != nil) {
        log.Fatal(err)
	}
	for _, file := range files {
		// fmt.Println(file); // 打印： &{src 114 2147484157 {970752721
        if (file.IsDir()) { // IsDir()：是否文件夹（目录），true：是目录，flase：非目录（是文件）
			fmt.Println("目录：", file.Name()) // FileInfo.Name() 返回文件名
		} else {
            fmt.Println("文件：", file.Name());
		}
	}

	/*
	b, err := ioutil.ReadFile("./abc.png"); 
	// func ReadFile(filename string) ([]byte, error) 读取文件中的所有数据，返回读取的数据(二进制流)和遇到的错误。 如果读取成功，则 err 返回 nil，而不是 EOF
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b) // 二进制流
    */

	//b, err := ioutil.ReadFile("/data/www/goStudy/hello.go")
	b, err := ioutil.ReadFile("./hello.go");
	if (err != nil) {
        log.Fatal(err)
	}
	str := string(b) // string([]byte) 把二进制流转为字符串
	fmt.Println(str)

	data := []byte("hell我\ngo\n"); // 字符串转二进制流
	// err := ioutil.WriteFile("test.txt", data, 0644) // 这样写都不行，为前面声明过err变量：no new variables on left side of :=
	e := ioutil.WriteFile("test.txt", data, 0644) 
	// func ioutil.WriteFile(filename string, data []byte, perm os.FileMode) error 写入二进制流数据入文件，写入前会清空文件。如果文件不存在则会以指定的权限创建它。
	check(e);

	// 像下面这种声明方式: f, err 虽然之前有声明过err，但是不会报错
	f, err := os.Open("test.txt"); // func os.Open(filename string) (*File, error) 打开一个文件，返回os.File结构体和错误信息
	check(err)
	b1 := make([]byte, 5) // 声明一个长度为5的二进制切片（数组）
	n1, err := f.Read(b1) // Read 从文件（句柄）中读取指定的字节数（一个字符一个字节），返回（读取到的）总字节数（整形）和错误信息
	check(err) // os.File.Read([]byte) 从文件的开头读取部分内容
	fmt.Printf("%d bytes:%s\n", n1, string(b1)) // 5 bytes:hell

	o2, err := f.Seek(4, 0) // os.File.Seek 从第几个字节开始读取（包括当前字节），这里指的是下标
	check(err)
	b2 := make([]byte, 5)
	n2, err := f.Read(b2) // 换行符也算在字符里面，占一个字符
	check(err)
	println(b2) // [5/5]0xc420055dbe   如果是 ： hell我\ngo\n 中文字，一个字占三个字节，如果从中间截取了，就无法转换成字符串
	fmt.Printf("%d bytes @ %d:%s\n", n2, o2, string(b2)) // 5 bytes @ 4:我\ng
	// n2, err := io.ReadAtLeast(f, b2, 5) // 这样写是会报错的
	// n3, err := io.ReadAtLeast(f, b2, 5)
	// check(err)
	// 上面两句这样写是会报错的： 原因是读取了文件的内容后，光标（游标）会再下一个字节中，所以文件读取遇到了结尾符EOF，不能再往下读取了
	/* panic: unexpected EOF

	goroutine 1 [running]:
	main.check(0x4db2e0, 0xc420086060)
		/data/www/go/src/goStudy/file.go:16 +0x4a
	main.main()
		/data/www/go/src/goStudy/file.go:105 +0x8d2
	exit status 2 */
	o3, err := f.Seek(4, 0)
	check(err)
	// num int := 3  // 这样声明是错误的，因为 := 是不允许声明类型的。 syntax error: unexpected int at end of statement
	num := 3
	b3 := make([]byte, num)
	n3, err := io.ReadAtLeast(f, b3, num) // io.ReadAtLeast(File,byte,int) 和上面的f.Read是一样的。读取文件的n个字节到一个变量
	check(err)
	fmt.Printf("%d bytes @ %d : %s\n", n3, o3, string(b3))

	index4, err := f.Seek(0, 0) // 移动读取位置到文件首地址
	check(err)
    fileInfo, _ := os.Stat("test.txt") // Stat：获取文件信息
	var fileSize int64 = fileInfo.Size() // Size : 获取文件大小。 返回的是int64而非int，如果赋值给int型的变量时，会报错
	num = int(fileSize) // num = 12 // 这个会报错，因为总共只有11个字节，遇到结束符报错：unexpected EOF
	byte4 := make([]byte, num)
	number4, err := io.ReadAtLeast(f, byte4, num)
	check(err)
    fmt.Printf("%d bytes @ %d : %s \n", number4, index4, string(byte4))

	f.Seek(0, 0)
	reader := bufio.NewReader(f) // 缓冲读取
	b4, err := reader.Peek(5) // 截取前5个字节
	check(err)
	fmt.Printf("5 bytes: %s \n", string(b4))

	f.Close() // 关闭文件句柄

	createTempDir()

	createTempFile()

	writeFile()

	file, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755);
	// OpenFile 打开一个文件，O_CREATE，如果不存在就创建，后面是创建的权限。O_RDONLY：只读方式打开
	if err != nil {
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	filePath := "./data2"
	startTime := time.Now() // 获取当前时间，精确到微妙（µs），1000毫秒（ms）=1s，1000µs=1ms
	readOne(filePath)
	endTime1 := time.Now()
	fmt.Printf("Cost time %v\n", endTime1.Sub(startTime))
	readTwo(filePath)
	endTime2 := time.Now()
	fmt.Printf("Cost time %v\n", endTime2.Sub(endTime1)) // 时间2 - 时间1
	nowTime := time.Now().Unix() // 转为linux时间戳（1970到现在的秒数）
	fmt.Printf("Now time %d seconds\n", nowTime)
}

// 演示创建临时文件夹
func createTempDir() {
	content := []byte("temporary file's content")  // 字符串转二进制流
	dirname, err := ioutil.TempDir("/tmp/Test", "example") // 创建临时目录
	// func TempDir(dir string, prefix string) (name string, err error) 在linux下的临时目录/tmp下创建一个目录。
	// dir为目录的绝对路径或相对（相对程序文件运行的）路径，如果为空（""）则默认在目录/tmp下；prefix为目录名的前缀。
	if (err != nil) {
        log.Fatal("创建临时文件夹失败：", err) // 如果指定的目录不存在或不可写时，将创建失败并返回错误信息
	}
	defer os.RemoveAll(dirname) // 删除目录（clean up），包括里面的所有文件。 defer ： 延迟执行，写在最后也是可以的，因为这句会再函数return后执行
    println("目录路径：", dirname); // 返回如：/tmp/example811899619
	tempFilePath := filepath.Join(dirname, "tempFile") // 连接两个文件路径
	println("临时文件路径：" + tempFilePath);
	if err := ioutil.WriteFile(tempFilePath, content, 0666); err != nil { // 将二进制流数据写入文件
		log.Fatal("写入流数据失败：", err)
	}
}

// 演示创建临时文件
func createTempFile() {
	content := []byte("临时的文件内容！！！！");
	tmpFile, err := ioutil.TempFile("", "test") // 在/tmp目录下创建一个临时文件，返回如：/tmp/test532523393
	// func TempFile(dir string, prefix string) (f *os.File, err error) 在dir目录中创建一个已prefix为前缀的文件，并将其以读写模式打开。返回创建的文件对象和遇到的错误。
	// dir为空时，默认在临时目录/tmp中创建文件，多次调用会创建不同的文件
	if (err != nil) {
        log.Fatal(err)
	}

	tmpFilePath := tmpFile.Name() // os.File.Name() 获取文件的完整路径，如：/tmp/test532523393
	defer os.Remove(tmpFilePath) // 删除文件（clean up）
	// log.Fatal("文件路径：" + tmpFilePath) // 特别注意：使用log.Fatal打印后会中断程序的进行，比较适合开发时用来断点调试
	println("文件路径：" + tmpFilePath);

    if length, err := tmpFile.Write(content); err != nil { // 写入二进制流文件，返回写入字节数和错误信息
		log.Fatal(err)
	} else {
        println("写入的字节数：", length)
	}
	if err := tmpFile.Close(); err != nil { // 关闭文件对象
        log.Fatal("关闭文件句柄出错：", err)
	}
}

// 演示写入文件
func writeFile() {
	file, err := os.Create("./data2") // Open a file for writing. 打开一个文件用于写入；打开文件时会清空里面的数据
	check(err)
	
	defer file.Close() // 关闭文件

	data := []byte{115, 111, 109, 101, 10, 97} // 声明一个二进制的切片并赋初始值； 10:\n  
	println("二进制数据[]byte{115, 111, 109, 101, 10, 97}打印为：", data)  
	number, err := file.Write(data) // Write ： 追加写入二进制数据
	check(err)
	fmt.Printf("Wrote %d bytes\n", number);
	number, err = file.Write([]byte("临时的文件内容！！！！")) 

	number, err = file.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", number)
	number, err = file.WriteString("你好我也好\n") // WriteString : 追加写入字符串数据

	file.Sync() // 貌似没有用，写不写都一样

	writer := bufio.NewWriter(file) // 声明一个缓冲写入对象
	number, err = writer.WriteString("buffered内容写入\n") // 写入缓冲区的字符串
	fmt.Printf("Wrote %d bytes\n", number)
	writer.Flush() // 清除缓冲区
}

// 下面演示两种常见的速度比较快的方式读取文件的内容

func readOne(path string) string {
	file, err := os.Open(path)
	if (err != nil) {
        panic(err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	return string(bytes)
}

func readTwo(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if (err != nil) {
        panic(err)
	}
    return string(bytes)
}
