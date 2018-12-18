// Go 常量
package main

import "fmt"

//  常量是一个简单值的标识符，在程序运行时，不会被修改的量。
//  常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。  array ,slice 或 map 不能声明为常量。
//  声明常量时可以不用指定数据类型，可以多个一起声明，且不能使用:=

const b string = "bb"
const status_normal, status_disable = 1, 0

const (
    UNKNOWN = 0
    FEMALE = 1
    MALE = 2
)

const (
    Unknown = 0
    Female = 1
    Male = 2
)

func main() {
    const LENGTH int = 10
    const WIDTH int = 5
    var area int
    area = LENGTH * WIDTH;
    
    fmt.Printf("面积为： %d\n", area);
    println(b, status_normal, status_disable); 
}

// iota，特殊常量，可以认为是一个可以被编译器修改的常量。
// 在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。