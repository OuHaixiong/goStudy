// Go 指针
// 和c语言一样，函数名本身就是指针
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

    t1 := t{"t1"}
    fmt.Println("m1调用前：", t1.Name) // m1调用前： t1
    t1.m1()
    fmt.Println("m1调用后：", t1.Name) // m1调用后： t1
    fmt.Println("m2调用前：", t1.Name) // m2调用前： t1
    t1.m2() // 这里相对于go自动进行了转换，把t1的地址传进去了
    fmt.Println("m2调用后：", t1.Name) // m2调用后： name2

    t2 := &t{"t2"} // 声明一个指针类型。  （不管是结构体变量还是指针结构体变量都拥有这两个方法）
    fmt.Println("m1调用前：", t2.Name) // m1调用前： t2
    t2.m1() // t2 是指针类型， 取 t2 的值并拷贝一份传给 m1。
    fmt.Println("m1调用后：", t2.Name) // m1调用后： t2
    fmt.Println("m2调用前：", t2.Name) // m2调用前： t2
    t2.m2() // 都是指针类型，不需要转换
    fmt.Println("m2调用后：", t2.Name) // m2调用后： name2

    t3 := t{Name:"t3"}
    fmt.Println("mm1调用前：", t3.Name) // mm1调用前： t3
    t3Name := mm1(t3)
    // t3Name := mm1(&t3) // 这个写法是不对的，因为需要传的参数为非指针
    fmt.Println("mm1调用后：", t3.Name) // mm1调用后： t3
    fmt.Println("mm1调用后，返回值为：", t3Name) // mm1调用后，返回值为： name1
    fmt.Println("mm2调用前：", t3.Name) // mm2调用前： t3
    mm2(&t3)
    // mm2(t3) // 这个写法是不对的，因为需要传的参数为指针
    fmt.Println("mm2调用后：", t3.Name) // mm2调用后： name1
    // 通过以上三例可以总结，调用函数时，参数类型一定需要传对，不然报错；但是如果是调用结构体的方法时，go会自动判断是否为指针，并转化传递

    // var t4 t = t{"t4"}
    // var i1 Interface1 = t4 // 不能这样写，这里会报错，因为对象或结构体需要实现接口的所有方法，而t没有实现m2()，t的指针实现了m2()。报错信息如下：
    // cannot use t4 (type t) as type Interface1 in assignment:
    // t does not implement Interface1 (m2 method has pointer receiver)
    // 但是可以如下这样写
    var t4 t = t{Name : "t4"}
    var i1 Interface1 = &t4 // 把t结构体的地址赋值给接口
    fmt.Println("使用接口地址，调用m2前：", t4.Name) // 使用接口地址，调用m2前： t4
    i1.m2() // 其实可以简单的理解一下接口，接口可以说是一种独特的结构体，相当于一个结构体的引用
    fmt.Println("使用接口地址，调用m2后：", t4.Name) // 使用接口地址，调用m2后： name2

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

type Interface1 interface {
    m1()
    m2()
}

type t struct {
    Name string
}
func (t1 t) m1() { // 可以简单理解为一个函数，传的是值拷贝
    t1.Name = "name1" // 所有t1属性的改变，并不会影响到包外部
}
// func (t1 t) m2() { // 这样写是不对的，报错：method redeclared: t.m2
// }
func (t2 *t) m2() { // 可以理解为一个函数，传的是引用
    t2.Name = "name2"
    return // 可以写return，也可以不写
}

func mm1(t1 t) string {
    t1.Name = "name1"
    return t1.Name
}

func mm2(t1 *t) {
    t1.Name = "name1" // 不管是值还是指针，访问结构体的属性均是这样写
}
