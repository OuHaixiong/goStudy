// Go 变量 . Go 语言变量名由字母、数字、下划线组成，其中首个字母不能为数字。 在func外部声明变量时不能使用:=
package main

var x, y int  // 未赋值的整型变量，默认为0。float32默认也是0，pointer默认nil
var(  // 类型不同时，多个全局变量的声明，这种只能出现在全局变量中，函数体内不支持
    a int
    b bool  // 布尔型，不赋值默认为false
)

var c, d int = 1, 2  // 多个一起赋值；如果声明类型，各变量的类型必须一致
var e, f = 123, "hello" // 也可以不声明类型，不声明是话go会自动判断

//这种不带声明格式的只能在函数体中出现
// g, h := 123, "hello" // syntax error: non-declaration statement outside function body

var i float32 = 3.1415 // 浮点型32位数据，

func main() { // := 为赋值操作符
    // println(h) // 所有的变量都是先声明后使用，如果使用了未声明的变量时，会报错： undefined: h
    // var bb string = "bb" // 同样的，如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误:bb declared and not used。 但是全局变量是允许声明但不使用。
    g, h := 123, "hello" // 省略了var的表示法。 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。
    var err string = "错误信息\n";
    print(err);
    println(x, y, a, b, c, d, e, f, g, h, i);
    println("a的内存地址是：", &a); // 字符串不能用单引号，使用后报错： invalid character literal (more than one character)
    
    var str string;
    str = `这就是一个字符串，还支持多行呢`;
    println(str);
}

// 指针:指的是这个变量的值是一个内存地址
// 如果你想要交换两个变量的值，则可以简单地使用 a, b = b, a。 
// 空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。 