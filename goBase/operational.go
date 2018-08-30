// Go 运算符（operational character）
package main

import "fmt"

// 下面演示位运算符：与&、或|、异或^、左移<<、右移>>、非^
/*
p   q  p&q p|q p^q
0   0   0   0   0
0   1   0   1   1
1   1   1   1   0
1   0   0   1   1
总结：左移n位就是乘以2的n次方；右移n位就是除以2的n次方
*/

func main() {
    var a uint = 60 // 0011 1100
    var b uint = 13 // 0000 1101
    var c uint = 0
    
    c = a & b // 12 = 0000 1100
    fmt.Printf("a & b 的值为 %d \n", c) 
    
    c = a | b // 61 = 0011 1101
    fmt.Printf("a | b 的值为 %d \n", c)
    
    c = a ^ b // 49 = 0011 0001
    fmt.Printf("a ^ b 的值为 %d \n", c);

    c = a << 2 // 240 = 1111 0000
    fmt.Printf("a << 2 的值为 %d \n", c);
    
    c = a >> 2 // 15 = 0000 1111
    fmt.Printf("a >> 2 的值为 %d \n", c);
    
//  c = ~a // Go语言不支持~符号（c语言中是~）
    c = ^a // 18446744073709551555 成了一个非常大的数，因为无符号整数是32位表示的，取反的话前面所有的0均变为1
    fmt.Printf("^a 的值为 %d \n", c);
    
    var x int = 3 // 注意了uint和int不能混用，出错提示：cannot assign int to a (type uint) in multiple assignment
    var y int = 5
    x, y = swap(x, y);
    fmt.Println("3和5交换后为：", x, y);
    var z int = 200
    z ^= 2
    fmt.Println("200 ^ 2 的值为：", z);
    
    tempFun();
}

/**
 * 通过异或的方式，快速的交换两个变量的值 （不需要使用第三个变量做中间变量）
 * (这种方式很高效)
 */
func swap(a, b int) (int, int) {
    a ^= b
    b ^= a
    a ^= b
    return a, b
}

func tempFun() {
    var a int = 4
    var b int32
    var c float32
    var ptr *int
    
    fmt.Printf("a变量的类型为：%T\n", a);
    fmt.Printf("b变量的类型为：%T\n", b);
    fmt.Printf("c变量的类型为： %T\n", c);
    
    ptr = & a // 也可以这样写&a；把a的地址赋值给ptr，ptr就成为一个指针类型的数据了
    fmt.Printf("a的值为：%d\n", a);
    fmt.Printf("*ptr的值为：%d\n", *ptr); // *ptr代码取内存地址上的值，指针可以简单理解为指向内存地址
    fmt.Print("ptr的值为:", ptr, "\n");  // 返回一个内存地址，类似：0xc420094068
}
