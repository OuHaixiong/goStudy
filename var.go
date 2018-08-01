// Go 变量
package main

var x, y int  // 未赋值的整型变量，默认为0
var( //这种只能出现在全局变量中，函数体内不支持
    a int
    b bool  // 布尔型，不赋值默认为false
)

var c, d int = 1, 2  // 多个一起赋值
var e, f = 123, "hello" // 也可以不声明类型

//这种不带声明格式的只能在函数体中出现
// g, h := 123, "hello"

var i float32 = 3.1415 // 浮点型32位数据，

func main() {
    g, h := 123, "hello" 
    println(x, y, a, b, c, d, e, f, g, h, i);
}
