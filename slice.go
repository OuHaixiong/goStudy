// Go 切片。 Go 语言切片（Slice）是对数组的抽象。Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),
// 与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大.说直白点就是：可变长数组
package main

import "fmt"

func main() {
    var number = make([]int, 3, 5); // 通过内置函数make([]Type, Length[, Capacity])声明一个切片，长度为3，容量为5;Length是数组的长度并且也是切片的初始长度
    // 最简单的声明是：var identifier(标识符) []Type 切片可以不声明长度（大小）; 另一种声明：var identifier []Type = make([]Type, len);可简写为：identifier := make([]Type, len) 
//    s := [] int {1, 2, 3}; // 直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
//    s := arr[:] // 初始化切片s,是数组arr的引用
//    s := arr[[startIndex]:[endIndex]] // 将arr中从下标startIndex到endIndex下的元素创建为一个新的切片，下标均可省略，缺省为第一个元素和最后一个元素
    fmt.Printf("len=%d cap=%d slice=%v\n", len(number), cap(number), number); // len=3 cap=5 slice=[0 0 0] ; len()获取长度，cap()求容量（切片最长可以达到多少）
    var numbers []int; // 空(nil)切片
    if (numbers == nil) { // 这里返回true
        fmt.Println("切片是空的");
    }
    
    numbers = [] int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // 对切片进行赋值
    printSlice(numbers);
    fmt.Println("numbers[1:4]=>", numbers[1:4]); // 返回：[1 2 3] 。 特别注意这里的下标（lower-bound下限）和上标（upper-bound上界：最大值）。 切片索引下标是包含的，上标是不包含的
    fmt.Println("numbers[:3]=>", numbers[:3]); // 返回：[0 1 2] 。 默认下标为0
    fmt.Println("numbers[5:]=>", numbers[5:]); // 返回：[5 6 7 8 9] 。 默认上标为最大值，len(numbers)=10
    
    number2 := numbers[:2] // len=2 cap=10 slice=[0 1]
    printSlice(number2);
    
    number3 := numbers[2:5] // len=3 cap=8 slice=[2 3 4]
    printSlice(number3);
}

func printSlice(x []int) { // int型的切片差数
    fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x);
}