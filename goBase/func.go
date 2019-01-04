// Go 函数。go中函数和方法不是一个概念，函数一般只具有一定功能的代码块，可直接调用，方法和函数类似，不同的是方法一般是挂在结构体中的
// 和c语言一样，函数名本身就是指针
package main;

import ( // 引用多个包时可以这样写
    "fmt"
    "math"
)
import "strings"
import "strconv"

func Version4ToInt(version string) int { // 函数的定义：func function_name([parameter list]) [return_types] {}
  arr := strings.Split(version, ".")
  fix_data := 100000000
  ver := 0
  for _, num := range arr {
    //d, error := strconv.Atoi(num) // error declared and not used（error不能用来做标识符，但err可以）
	d, _error := strconv.Atoi(num) // 字符串转整型（函数有两个返回值，）
    if _error != nil { // 单个_表示不接受参数(该值被直接抛弃掉)，且_不能用来作为变量（cannot use _ as value）
	    fmt.Println("字符串转换成整数失败")
	}
	ver += d*fix_data
    fix_data = fix_data / 100
  }
  return ver
}

func main() { // 如果是需要go来运行的一定不能少了main函数，除非你的文件是被包含进来的 （Go 语言最少有个 main() 函数）
    var version string = "1.1.0";
    var ver int = Version4ToInt(version);
    //fmt.Println(version, ver); 
	println(version, ver); // 两点需要注意的：
	// 1，打印一定需要在函数里面，不能在外面： non-declaration statement outside function body
	// 2，打印是需要引入fmt包的，如果你不用fmt.Println而直接用println，其实是go帮你自动引入了包fmt
	
	var a int = 10;
//	a int = 10; // 局部声明变量也不能少var，不然报错：syntax error: unexpected int at end of statement
	var b int = 15;
	var maxNumber int = max(a, b);
	println(maxNumber);
	
	var x = "Hello";
	var y = "world";
	println(x, y);
    //println(swapString(x, y)); // 不能直接写在打印函数中，会报错：multiple-value swapString() in single-value context（有多个返回值是不行的，一个是可以的）
    x, y = swapString(x, y);
    println(x, y);
    
    // x = 100; // 这样是不行的，之前的类型是string，现在又是int，在go中不允许，报错：cannot use 100 (type int) as type string in assignment
    p, q := 100, 200; // 不声明类型的话，赋值整形的话，其类型为int
    println("交换前p, q的值为：", p, q);
    swapInt(&p, &q); // 按引用（地址）传递时，需要这样写。
    // &p 指向 p 指针，p 变量的地址
    // &q 指向 q 指针，q 变量的地址
    println("交换后p, q的值为：", p, q);
    
    getSquareRoot := func(x float64) float64 { return math.Sqrt(x);}; // 声明函数变量
    // 上面的这种形式的写法就和JavaScript的写法是一样的，貌似php也是可以这样写的
    println(getSquareRoot(9));
    
    nextNumber := getSequence(); // i=0  nextNumber 为一个函数
    fmt.Println(nextNumber()); // i=1
    fmt.Println(nextNumber()); // i=2
    fmt.Println(nextNumber()); // i=3
    nextNumber2 := getSequence(); // i=0
    fmt.Println(nextNumber2()); // i=1
    fmt.Println(nextNumber2()); // i=2
    
    
    var c Circle; // 定义一个变量为结构体
    c.radius = 10.00; // 给结构体的属性赋值
    fmt.Println("Area of Circle(c) = ", c.getArea()); // 调用结构体的方法
}

func max(num1, num2 int) int {
    if (num1 > num2) {
        return num1;
    } else {
        return num2;
    }
}

// 下面演示函数的多个返回值
func swapString(x, y string) (string, string) {
    return y, x; // 交换两个字符串
}

// 下面演示按引用传递参数
func swapInt(x, y *int) { // 这样写是可以的，但是推荐写法如：func swapInt(x *int, y *int)
    *x ^= *y; // 引用地址，通常也是我们说的指针类型，如果需要取地址（指针）中的值，需要用 *
    *y ^= *x;
    *x ^= *y;
}

// 下面演示匿名函数和闭包的作用
func getSequence() func() int {
    i := 0
    return func() int { // 匿名函数，也算是一个闭包
        i += 1;
        return i;
    }
} 

// 定义一个结构体类型，结构体就类似php中的类，但是只有公共属性，方法可以动态赋予
type Circle struct { // 圆的结构体
    radius float64  // 半径
}

// 给结构体定义一个方法 (类似php中类的一个方法)
func (c Circle) getArea() float64 { // 该 method 属于 Circle 类型对象中的方法
    return 3.14 * c.radius * c.radius;
}


// go和php中函数的不同：
// go中的函数可以有两个或以上的返回值，但php中只能有一个返回值
// php中返回值类型可以不声明，但go中一定要声明

