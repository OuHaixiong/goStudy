package main;
var a = "欧海雄第一次学习go开发";
var b string = "字符型数据"; // 字符串一定是双引号，不能使用单引号。Go语言的字符串的字节使用UTF-8编码标识Unicode文本
var c bool; // 布尔型的数据，如果没有赋值，默认就是false

func main() {
    println(a, b, c)
}