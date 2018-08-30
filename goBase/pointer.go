// Go 指针
package main

import (
    "fmt"
    
    
)

const MAX_LENGTH int = 3;

func main() {
    var a int = 10;
    fmt.Printf("变量a的地址：%x \n", &a); // c420094010
    
    var x int = 20; // 声明一个整形变量
    var ip * int; // 声明一个整形的指针，指针在使用前也是需要先声明的
    ip = &x; // 指针指向变量的存储地址
    
    fmt.Printf("x变量的地址是：%x\n", &x);
    fmt.Printf("ip(指针)变量的存储地址是：%x\n", ip);
    fmt.Printf("ip变量的值（*ip）是：%d\n", *ip); // *pointer 用于取指针的值
    
    var ptr *int;
    fmt.Printf("ptr的值为：%x\n", ptr); // 这里返回0，是一个空指针
    if (ptr == nil) {
        fmt.Println("ptr是空指针");
    } else {
        fmt.Println("ptr不是空指针");
    }
    
    pointerArray();
    
    pointerPointer();
}

// 下面演示指针数组
func pointerArray() {
    a := [MAX_LENGTH]int{10, 100, 200};
    var i int;
    var ptr [MAX_LENGTH]*int;
    for i=0; i<MAX_LENGTH; i++ {
        ptr[i] = & a[i]; // 整形地址赋值给指针数组
    }
    for i=0; i<MAX_LENGTH; i++ {
        fmt.Printf("a[%d] => %d(%x)\n", i, *ptr[i], ptr[i]);
    }
}

func pointerPointer() { // 指向指针的指针：如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
    var a int;
    var ptr *int;
    var pptr **int;
    a = 3000;
    ptr = &a; // 指针赋值，指针ptr指向a变量的地址
    pptr = &ptr; // 指针赋值，指针pptr指向ptr指针变量的地址
    
    fmt.Printf("变量a = %d\n", a);
    fmt.Printf("指针变量 *ptr = %d\n", *ptr);
    fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr);
}