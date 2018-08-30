package abc

import (
	"fmt"
)


func Wl(str string) { // 大写字母开头的函数代表公有函数，可以被外部调用（is exported）
    fmt.Println("wuliao.go Wl function print:", str);
}

func wor() { // 小写字母开头的代表私有的函数，不能被外部调用（unexported）
	const LENGTH int = 10
	const WIDTH int = 5
    println("wuliao.go wor function print const:", LENGTH, WIDTH);
}

// func kao() { // 在同一个包里的函数是不允许重复的，不管是私有还是公有的
	
// }

func init() {
	println("这里是wuliao的init");
	fmt.Println("wuliao.go const ‘BB’:", BB);
	println("wuliao.go variable:", fullName, age);
}

var fullName, age = "欧阳海雄", 35
const BB string = "bb" // 常量声明在后面也是可以的

// fmt.Println("wuliao.go const ‘BB’:", BB); // 打印一定不能在函数外面
